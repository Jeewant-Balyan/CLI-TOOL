# DBMASTER – QuickDB & Database Utility CLI Tool  

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/Jeewant-Balyan/CLI-TOOL)  

DBMASTER is a **Command Line Interface tool** designed for students, startups, and hackathons to **quickly create, visualize, seed, and export databases**. It supports **QuickDB in-memory SQLite**, random data generation, ER diagram visualization, multi-format exports, and a mini dashboard for easy demos.  

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
[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/Jeewant-Balyan/CLI-TOOL)  

Click the badge above to launch your project in **GitHub Codespaces**.  

### 2. Update & Install Dependencies  

Run these inside Codespaces terminal:  

```bash
# Update packages
sudo apt update && sudo apt upgrade -y  

# Install Go, SQLite, and Graphviz
sudo apt install -y golang-go sqlite3 graphviz unzip git build-essential  

# Optional: Install Excel export dependency
go install github.com/xuri/excelize/v2@latest  
```  

### 3. Build the CLI Tool  

```bash
# Fix the permission issue
chmod +x dbmaster  

# Build the Go project
go build -o dbmaster main.go  
```  

---

## ⚡ Usage Guide  

### 1. QuickDB – Launch in-memory database  

```bash
./dbmaster quickdb
```

**This will:**  
- Spin up an in-memory SQLite DB.  
- Auto-create tables (users, orders).  
- Seed with random sample data.  
- Let you run stats, ER diagram, exports, etc.  

**Output:**  

```
QuickDB in-memory SQLite running.
Press Enter to stop QuickDB...
```  

---

### 2. Generate ER Diagram  

```bash
./dbmaster visualize
```

```
ER diagram generated as schema.png
```  

---

### 3. Export Table Data  

**CSV:**  
```bash
./dbmaster export users csv users.csv
```
```
Exported users as csv to users.csv
```  

**JSON:**  
```bash
./dbmaster export users json users.json
```
```
Exported users as json to users.json
```  

**Excel:**  
```bash
./dbmaster export users excel users.xlsx
```
```
Usage: dbmaster export <table> <format> <outfile>
```  

---

### 4. Mini Dashboard  

```bash
./dbmaster dashboard
```

Outputs table stats, ER diagram, and export package ZIP.  

---

### 5. Verify Template SQL File  

```bash
ls -R | grep template_ecom.sql
```
```
template_ecom.sql
```  

---

## 🗂️ Project Structure  

```
CLI-TOOL/
│── main.go
│── dbmaster (built binary)
│── template_ecom.sql
│── template_school.sql
│── template_library.sql
│── export_package.zip (after running dashboard)
```

---

## 💡 Recommended Workflow for Hackathon Demo  

1. Launch QuickDB:  
```bash
./dbmaster quickdb
```

2. Load template DB (e.g., E-commerce):  
```bash
./dbmaster template ecom
```

3. Seed some sample data:  
```bash
./dbmaster quickseed --rows=20
```

4. Visualize ER diagram:  
```bash
./dbmaster visualize
```

5. Export tables:  
```bash
./dbmaster export users csv users.csv
./dbmaster export orders json orders.json
```

6. Launch dashboard:  
```bash
./dbmaster dashboard
```

7. Export package:  
```bash
./dbmaster package
```

---

## 📸 Media (Screenshots & Video Demo)  

**Video Demo:**  
* [ ] Add link to Codespaces demo video  

**Screenshots:**  
![WhatsApp Image 2025-09-29 at 22 51 14_98cf30f0](https://github.com/user-attachments/assets/6e36bc5d-e546-415e-8593-0723d6f87f69)

---

## 🖥️ Setup on Local Windows Machine  

### 1. Install Dependencies  

1. **Install Go:**  
   Download from [https://go.dev/dl/](https://go.dev/dl/) and follow installer instructions.  
2. **Install SQLite:**  
   Download from [https://sqlite.org/download.html](https://sqlite.org/download.html) and add to PATH.  
3. **Install Graphviz:**  
   Download from [https://graphviz.org/download/](https://graphviz.org/download/) and add to PATH.  
4. **Optional: Excel export dependency**  
   Open PowerShell or terminal and run:  
   ```powershell
   go install github.com/xuri/excelize/v2@latest
   ```  

### 2. Open Terminal  

Open **PowerShell** or **Command Prompt** in the project folder where `main.go` exists.  

### 3. Build DBMASTER  

```powershell
# Build the Go project
go build -o dbmaster.exe main.go
```

If you get a permission issue, you can run as Administrator.  

### 4. Run DBMASTER  

```powershell
.\dbmaster.exe quickdb
```

Other commands are the same as Codespaces:  

```powershell
.\dbmaster.exe visualize
.\dbmaster.exe export users csv users.csv
.\dbmaster.exe dashboard
```  

> 💡 Tip: Use `.\dbmaster.exe` instead of `./dbmaster` on Windows.

## 👥 Team Information  

| Name              | GitHub |  
| ----------------- | ---------------------------- |  
| Shubham Chaudhary | [shubhamchaudhary-dev](https://github.com/shubhamchaudhary-dev) |  
| Jeewant Balyan    | [jeewant-balyan](https://github.com/jeewant-balyan) |  
