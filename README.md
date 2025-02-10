Here’s your **README.md** with your GitHub link included.  

---

# **Excel to JSON Converter (Go)**

This Go application converts an **Excel sheet** to a **structured JSON file**.  
Users can specify the **Excel file path**, **sheet name**, and **output JSON file name** via command-line flags.

---

## **Installation**

1. **Install Go** (if not installed):  
   Download and install Go from [golang.org](https://go.dev/dl/).

2. **Clone the Repository**:
   ```sh
   git clone https://github.com/4hmedhabib/excelToJson.git
   cd excelToJson
   ```

3. **Install Dependencies**:  
   ```sh
   go get github.com/xuri/excelize/v2
   ```

---

## **Usage**

### **Run the Application**
```sh
go run main.go -path sample.xlsx -sheet Sheet1 -output output.json
```

### **Command-Line Flags**
| Flag      | Description                     | Example                 |
|-----------|---------------------------------|-------------------------|
| `-path`   | Path to the Excel file         | `-path sample.xlsx`     |
| `-sheet`  | Name of the sheet to convert   | `-sheet Sheet1`         |
| `-output` | Output JSON file name          | `-output output.json`   |

---

## **Example**

### **Input Excel File (`sample.xlsx`)**
#### **Sheet: `Sheet1`**
| Name  | Age |
|--------|----|
| John   | 30  |
| Alice  | 25  |

### **Run Command**
```sh
go run main.go -path sample.xlsx -sheet Sheet1 -output output.json
```

### **Generated JSON (`output.json`)**
```json
[
  {
    "Name": "Sinwar",
    "Age": "30"
  },
  {
    "Name": "Dayf",
    "Age": "25"
  }
]
```

---

## **Building an Executable**
To build a standalone executable:
```sh
go build -o excelToJson
```
Run the compiled binary:
```sh
./excelToJson -path sample.xlsx -sheet Sheet1 -output output.json
```

---

## **Error Handling**
- If the **Excel file does not exist**, the program will return an error.
- If the **specified sheet is missing**, the program will inform the user.
- If **no command-line arguments are provided**, the program will show a usage guide.

---

## **License**
This project is open-source and free to use under the **MIT License**.

---

## **Author**
👤 **Ahmed Habib**   
🔗 [GitHub](https://github.com/4hmedhabib/excelToJson)

--- 