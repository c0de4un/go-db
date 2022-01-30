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
  
