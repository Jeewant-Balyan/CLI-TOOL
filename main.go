package main

import (
	"archive/zip"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xuri/excelize/v2"
)

type TableInfo struct {
	Name    string
	Columns []ColumnInfo
}

type ColumnInfo struct {
	Name string
	Type string
	Null bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dbcli <command> [options]")
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "quickdb":
		runQuickDB()
	case "dashboard":
		runDashboard()
	case "visualize":
		runVisualize()
	case "export":
		runExport()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func runQuickDB() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("QuickDB in-memory SQLite running.")
	loadTemplate(db)
	seedRandomData(db)
	fmt.Println("Press Enter to stop QuickDB...")
	fmt.Scanln()
}

func loadTemplate(db *sql.DB) {
	files := []string{"template_ecom.sql", "template_school.sql"}
	for _, f := range files {
		if _, err := os.Stat(f); err == nil {
			sqlBytes, _ := ioutil.ReadFile(f)
			db.Exec(string(sqlBytes))
		}
	}
}

func seedRandomData(db *sql.DB) {
	rand.Seed(time.Now().UnixNano())
	db.Exec("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY,name TEXT,email TEXT)")
	db.Exec("CREATE TABLE IF NOT EXISTS orders(id INTEGER PRIMARY KEY,user_id INTEGER,total INTEGER)")
	for i := 1; i <= 5; i++ {
		name := fmt.Sprintf("User%d", i)
		email := fmt.Sprintf("user%d@example.com", i)
		db.Exec("INSERT INTO users(name,email) VALUES(?,?)", name, email)
		total := rand.Intn(500) + 50
		db.Exec("INSERT INTO orders(user_id,total) VALUES(?,?)", i, total)
	}
}

func runDashboard() {
	fmt.Println("=== Mini DB Dashboard ===")
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	loadTemplate(db)
	seedRandomData(db)
	showStats(db)
	fmt.Println("ER Diagram will be generated as schema.png")
	runVisualize()
	fmt.Println("Exporting combined ZIP of diagram + CSV...")
	exportAll(db)
	fmt.Println("Dashboard demo complete.")
}

func showStats(db *sql.DB) {
	tables := getTables(db)
	for _, t := range tables {
		count := 0
		db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", t)).Scan(&count)
		fmt.Printf("Table: %s, Rows: %d\n", t, count)
		cols := getColumns(db, t)
		for _, c := range cols {
			nulls := 0
			db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s IS NULL", t, c.Name)).Scan(&nulls)
			fmt.Printf("  Col: %s (%s), Nulls: %d\n", c.Name, c.Type, nulls)
		}
	}
}

func getTables(db *sql.DB) []string {
	rows, _ := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	defer rows.Close()
	var tables []string
	for rows.Next() {
		var t string
		rows.Scan(&t)
		tables = append(tables, t)
	}
	return tables
}

func getColumns(db *sql.DB, table string) []ColumnInfo {
	rows, _ := db.Query(fmt.Sprintf("PRAGMA table_info(%s)", table))
	defer rows.Close()
	var cols []ColumnInfo
	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull, pk int
		var dflt interface{}
		rows.Scan(&cid, &name, &ctype, &notnull, &dflt, &pk)
		cols = append(cols, ColumnInfo{name, ctype, notnull == 0})
	}
	return cols
}

func runVisualize() {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()
	loadTemplate(db)
	seedRandomData(db)
	tables := getTables(db)
	var dot strings.Builder
	dot.WriteString("digraph G {\nnode [shape=box style=filled fillcolor=lightblue];\n")
	for _, t := range tables {
		cols := getColumns(db, t)
		label := t + "\\n"
		for _, c := range cols {
			label += c.Name + ":" + c.Type + "\\n"
		}
		dot.WriteString(fmt.Sprintf("\"%s\" [label=\"%s\"];\n", t, label))
	}
	dot.WriteString("}\n")
	ioutil.WriteFile("schema.dot", []byte(dot.String()), 0644)
	exec.Command("dot", "-Tpng", "schema.dot", "-o", "schema.png").Run()
	fmt.Println("ER diagram generated as schema.png")
}

func runExport() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: dbcli export <table> <format> <outfile>")
		return
	}
	table := os.Args[2]
	format := os.Args[3]
	outfile := os.Args[4]
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()
	loadTemplate(db)
	seedRandomData(db)
	rows, _ := db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	defer rows.Close()
	cols, _ := rows.Columns()
	values := make([]interface{}, len(cols))
	valuePtrs := make([]interface{}, len(cols))
	for i := range values {
		valuePtrs[i] = &values[i]
	}
	switch format {
	case "csv":
		f, _ := os.Create(outfile)
		defer f.Close()
		w := csv.NewWriter(f)
		w.Write(cols)
		for rows.Next() {
			rows.Scan(valuePtrs...)
			var rec []string
			for _, v := range values {
				rec = append(rec, fmt.Sprintf("%v", v))
			}
			w.Write(rec)
		}
		w.Flush()
	case "json":
		var arr []map[string]interface{}
		for rows.Next() {
			rows.Scan(valuePtrs...)
			m := map[string]interface{}{}
			for i, c := range cols {
				m[c] = values[i]
			}
			arr = append(arr, m)
		}
		b, _ := json.MarshalIndent(arr, "", "  ")
		ioutil.WriteFile(outfile, b, 0644)
	case "excel":
		f := excelize.NewFile()
		sheet := "Sheet1"
		for i, c := range cols {
			f.SetCellValue(sheet, fmt.Sprintf("%s1", string('A'+i)), c)
		}
		rowNum := 2
		for rows.Next() {
			rows.Scan(valuePtrs...)
			for i, v := range values {
				f.SetCellValue(sheet, fmt.Sprintf("%s%d", string('A'+i), rowNum), v)
			}
			rowNum++
		}
		f.SaveAs(outfile)
	}
	fmt.Println("Exported", table, "as", format, "to", outfile)
}

func exportAll(db *sql.DB) {
	tables := getTables(db)
	tmpDir := "tmp_export"
	os.Mkdir(tmpDir, 0755)
	for _, t := range tables {
		out := fmt.Sprintf("%s/%s.csv", tmpDir, t)
		os.Args = []string{"dbcli", "export", t, "csv", out}
		runExport()
	}
	os.Rename("schema.png", tmpDir+"/schema.png")
	zipFile, _ := os.Create("export_package.zip")
	defer zipFile.Close()
	zipw := zip.NewWriter(zipFile)
	defer zipw.Close()
	filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f, _ := os.Open(path)
			defer f.Close()
			w, _ := zipw.Create(info.Name())
			io.Copy(w, f)
		}
		return nil
	})
	fmt.Println("Combined export package created as export_package.zip")
}
