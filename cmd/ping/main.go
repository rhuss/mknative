package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    target := os.Getenv("TARGET")
    if target == "" {
        log.Fatal("missing environment variable TARGET")
    }
    log.Printf("target:  %s\n", target)

    transport := http.DefaultTransport.(*http.Transport).Clone()

    // make sure to create a connection per request
    transport.DisableKeepAlives = true
    transport.MaxIdleConnsPerHost = -1
    client := http.Client{
        Timeout: 2 * time.Second,
        Transport: transport,
    }

    for {
        resp, err := client.Get(target)
        if err != nil {
            log.Printf("ping failed: %v", err)
        } else {
            if resp.StatusCode == 200 {
                log.Println("ping succeeded")
            } else {
                payload, err := io.ReadAll(resp.Body)
                if err != nil {
                    log.Printf("an error occurred while reading the HTTP response: %v\n", err)
                } else {
                    resp.Body.Close()
                    log.Printf("ping failed: %s\n", string(payload))
                }
            }
        }

        time.Sleep(2 * time.Second)

    }
}
