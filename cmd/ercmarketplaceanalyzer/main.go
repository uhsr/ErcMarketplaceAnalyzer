// cmd/ercmarketplaceanalyzer/main.go
package main

import (
"flag"
"log"
"os"

"ercmarketplaceanalyzer/internal/ercmarketplaceanalyzer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ercmarketplaceanalyzer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
