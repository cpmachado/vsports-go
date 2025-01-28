# VSports API in golang

## Implements most methods of api

### Sample use

```go
package main

import (
 "github.com/sapo/vsports-go/client"
 "fmt"
 "time"
)

func main() {
 fmt.Println("Demo of VSports API Go.")

 // If you want to add logging to the client, uncomment these lines
 // And add imports for "log/slog" and "os"
 // loggerClientHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
 // loggerClient := slog.New(loggerClientHandler).With("source", "VSPORTS-API")

 var config = client.ClientConfig{
  APIKey:         "my_api_key",
  TimeoutSeconds: 10,
  RedisConfig: client.RedisConfig{
   Addr:     "localhost:6379",
   Password: "",
   DB:       0,
  },
  CacheDuration: 5, // 5 minutes
 }

 // Create the client
 // Optionally, you can pass a logger object to the client (see above)
 client, err := client.VSportsClient(config, nil)
 if err != nil {
  fmt.Printf("Error creating client: %v", err)
  return
 }

 // Get all events for today
 today := time.Now().Format("2006-01-02")
 events, err := client.GetEventsByDate(today, today, true)

 if err != nil {
  fmt.Printf("Error getting events: %v", err)
 } else {
  if len(events) == 0 {
   fmt.Println("No events found for today. Try another date.")
  } else {
   fmt.Printf("Total number of events found for today: %d\n", len(events))
   for _, event := range events {
    fmt.Printf("Event: %d: %s x %s at %s\n", event.ID, event.TeamA.Name, event.TeamB.Name, event.DateTime)
   }
  }
 }
}

```
