package main
import (
    "log"
    "os"
    "[github.com/pocketbase/pocketbase](https://github.com/pocketbase/pocketbase)"
    "[github.com/pocketbase/pocketbase/core](https://github.com/pocketbase/pocketbase/core)"
)
func main() {
    app := pocketbase.New()
    app.OnAfterBootstrap().PreAdd(func(e *core.BootstrapEvent) error {
        log.Println("----------------------------------------")
        log.Println("  Project VISION AI Backend is starting...")
        log.Println("----------------------------------------")
        return nil
    })
    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
