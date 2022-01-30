# database
Golang based database realization

# Description
Example project of database realization using drivers and models

# Requirements
 * Go 1.7+;

# Testign
 * run: <go test>
 
# Examples:
## Connect/Disconnect
```go
config := database.Configuration
config.SetUri("")
config.SetDatabaseName("myDB")
config.SetLogin("")
config.SetPassword("")
config.SetEncoding(database.ENCODING_UTF8)

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
  
