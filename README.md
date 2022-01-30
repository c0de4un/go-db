# database
Golang based database realization

# Description
Example project of database realization using drivers and models

# Requirements
 * Go 1.17+;

# Test
 * $go test

# Examples:
## Connect/Disconnect
```go
config := database.LoadConfig("database.cfg")

driver, err := database.NewDriver(database.FILE_DRIVER)
if err != nil {
  log.Fatal(err)
}

err = driver.Configure(config)
if err != nil {
  log.Fatal(err)
}

connection, err := driver.Connect()
if err != nil {
  log.Fatal(err)
}

// Read    . . .
// Iterate . . .
// Write   . . .

defer connection.Disconnect()
```

## Read/Write
```go
// var driver database.IDriver (FileDriver)
// var query database.IQuery (FileQuery)
defer driver.Disconnect()

row := driver.ReadRow("users", query)

row.SetCol("name", "John")

driver.WriteRow(row)
```

## Building a Query
```go
// var driver database.IDriver (FileDriver)
// var query database::IQuery
query := driver.NewQuery()

query.Select("users", "*")
query.From("users")
query.Where("users", "id", "=", "1")
// var row database.Row
row := query.GetRow()

fmt.Println("user name is ", row.GetCol("name"))
```
