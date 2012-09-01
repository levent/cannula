package main

import (
  "net"
  "bufio"
  "os"
  "fmt"
)

func main() {
  listen := "api.cosm.com:8081"
  conn, error := net.Dial("tcp", listen)
  if error != nil {
    fmt.Printf("Cannot dial: %s\n", error)
    os.Exit(1)
  }
  api_key := os.Getenv("API_KEY")
  feed_id := os.Getenv("FEED_ID")
  json := fmt.Sprintf("{\"method\":\"get\",\"resource\":\"/feeds/%s\",\"headers\":{\"X-ApiKey\":\"%s\"}}", feed_id, api_key)
  fmt.Printf("Sent:\n%s\n\n", json)
  fmt.Fprintf(conn, json)
  status, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Printf("Received:\n%s\n\n", status)
}
