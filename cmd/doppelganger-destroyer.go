package main

import (
    "flag"
    "fmt"
    "log"

    "doppelganger-destroyer/internal/doppelganger"
)

func main() {
    rootDirPtr := flag.String("dir", ".", "Root directory to search for duplicates")
    helpPtr := flag.Bool("help", false, "Show help message")
    flag.Parse()

    if *helpPtr {
        fmt.Println("Usage:")
        fmt.Println("  -dir string: Root directory to search for duplicates (default: current directory)")
        fmt.Println("  -help: Show this help message")
        return
    }

    if err := doppelganger.FindDuplicates(*rootDirPtr); err != nil {
        log.Fatal(err)
    }
}
