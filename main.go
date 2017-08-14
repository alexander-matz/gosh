package main;

import (
    "os"
    "log"
    "net/http"
    "flag"
    "fmt"
)

var usage = func() {
    fmt.Fprintf(os.Stderr, "Simple static file server with no dependencies\n")
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Usage:\n")
    fmt.Fprintf(os.Stderr, "  gosh [options]\n")
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Options:\n")
    flag.PrintDefaults()
}

func main() {
    flag.Usage = usage
    var dir = flag.String("dir", ".", "serve files from directory")
    var addr = flag.String("addr", "127.0.0.1:3000", "binding address")
    var help = flag.Bool("help", false, "show this help")
    flag.Parse()
    if (*help) {
        usage()
        os.Exit(0)
    }

    fs := http.FileServer(http.Dir(*dir))
    http.Handle("/", fs)

    log.Println("Listening...")
    http.ListenAndServe(*addr, nil)
}
