# DBCLI – QuickDB & Database Utility CLI Tool

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/Jeewant-Balyan/CLI-TOOL)

DBCLI is a **Command Line Interface tool** designed for students, startups, and hackathons to **quickly create, visualize, seed, and export databases**. It supports **QuickDB in-memory SQLite**, random data generation, ER diagram visualization, multi-format exports, and a mini dashboard for easy demos.

---

## 🌟 Features

* **QuickDB**: Launch an in-memory SQLite database for instant testing or demos.
* **Quick Seed**: Generate fake users, orders, or any sample data automatically.
* **DB Stats & Summary**: View row counts, null percentages, and column data types.
* **Interactive/Colored ER Diagrams**: Generate schema diagrams highlighting relationships and table types.
* **Multi-Format Export**: Export tables as CSV, JSON, Excel, or a combined ZIP.
* **Mini Dashboard**: Combine stats, ER diagram, and exports in one CLI view.
* **Quick Template DBs**: Pre-built templates like E-commerce, School, or Library.

---

## 🛠️ Setup in GitHub Codespaces

### 1. Open in Codespaces

Click the badge above to launch your project in **GitHub Codespaces**.

### 2. Install Dependencies

Inside the Codespaces terminal, run:

```bash
sudo apt update && sudo apt install -y sqlite3 graphviz
```

### 3. Build the CLI Tool

```bash
# Fix the permission issue
chmod +x dbcli

# Build the Go project
go build -o dbcli main.go
```

---

## ⚡ Usage Guide

### 1. QuickDB – Launch in-memory database

```bash
./dbcli quickdb
```

**Output:**

```
QuickDB in-memory SQLite running.
Press Enter to stop QuickDB...
```

### 2. Generate ER Diagram

```bash
./dbcli visualize
```

**Output:**

```
ER diagram generated as schema.png
```

### 3. Export Table Data

**CSV:**

```bash
./dbcli export users csv users.csv
```

**Output:**

```
Exported users as csv to users.csv
```

**JSON:**

```bash
./dbcli export users json users.json
```

**Output:**

```
Exported users as json to users.json
```

**Excel:**

```bash
./dbcli export users excel users.xlsx
```

**Output:**

```
Usage: dbcli export <table> <format> <outfile>
```

### 4. Mini Dashboard

```bash
./dbcli dashboard
```

**Sample Output:**

```
=== Mini DB Dashboard ===
Table: products, Rows: 0
  Col: id (INTEGER), Nulls: 0
  Col: name (TEXT), Nulls: 0
  Col: price (INTEGER), Nulls: 0
...
ER Diagram will be generated as schema.png
ER diagram generated as schema.png
Exporting combined ZIP of diagram + CSV...
Exported products as csv to tmp_export/products.csv
...
Combined export package created as export_package.zip
Dashboard demo complete.
```

### 5. Verify Template SQL File

```bash
ls -R | grep template_ecom.sql
```

**Output:**

```
template_ecom.sql
```

---

## 🗂️ Project Structure

```
CLI-TOOL/
│── main.go
│── dbcli (built binary)
│── template_ecom.sql
│── template_school.sql
│── template_library.sql
│── export_package.zip (after running dashboard)
```

---

## 💡 Recommended Workflow for Hackathon Demo

1. Launch QuickDB:

```bash
./dbcli quickdb
```

2. Load a template DB (e.g., E-commerce):

```bash
./dbcli template ecom
```

3. Seed some sample data:

```bash
./dbcli quickseed --rows=20
```

4. Visualize ER diagram:

```bash
./dbcli visualize
```

5. Export tables:

```bash
./dbcli export users csv users.csv
./dbcli export orders json orders.json
```

6. Launch the mini dashboard:

```bash
./dbcli dashboard
```

7. Combined export:

```bash
./dbcli package
```

---

## 📸 Media (Screenshots & Video Demo)

**Screenshots:**

* [ ] Add screenshot of QuickDB running
* [ ] Add screenshot of ER diagram `schema.png`
* [ ] Add screenshot of CSV/Excel export

**Video Demo:**

* [ ] Add link or embed video showing end-to-end demo in Codespaces

---

## 👥 Team Information

| Name | Role | Contact |
| ---- | ---- | ------- |
|      |      |         |
|      |      |         |
|      |      |         |

