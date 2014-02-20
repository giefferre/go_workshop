package main

import (
    "io"
    "log"
    "net"
    "fmt"
)

const listenAddr = "localhost:4000"
var partner = make(chan io.ReadWriteCloser)

func chat(a, b io.ReadWriteCloser) {
    fmt.Fprintln(a, "Found one! Say hi.")
    fmt.Fprintln(b, "Found one! Say hi.")

    // It's important to clean up when the conversation is over.
    // To do this we send the error value from each io.Copy call to a channel,
    //  log any non-nil errors, and close both connections. 
    errc := make(chan error, 1)
    go cp(a, b, errc)
    go cp(b, a, errc)
    if err := <-errc; err != nil {
        log.Println(err)
    }
    a.Close()
    b.Close()
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
    _, err := io.Copy(w, r)
    errc <- err
}

func match(c io.ReadWriteCloser) {
    fmt.Fprint(c, "Waiting for a partner...")
    select {
    case partner <- c:
        // now handled by the other goroutine
    case p := <-partner:
        chat(p, c)
    }
}


func main() {
    l, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatal(err)
    }

    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go match(c)
    }
}

/*

HTTP + WEBSOCKET VERSION

package main

import (
    "io"
    "log"
    "fmt"
    "net/http"
    "html/template"

    "code.google.com/p/go.net/websocket"
)

const listenAddr = "localhost:4000"
var partner = make(chan io.ReadWriteCloser)

type socket struct {
    io.ReadWriter
    done chan bool
}

func (s socket) Close() error {
    s.done <- true
    return nil
}

func socketHandler(ws *websocket.Conn) {
    s := socket{ws, make(chan bool)}
    go match(s)
    <-s.done
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("./template.html")

    if err != nil {
        log.Fatal(err)
    }

    t.Execute(w, listenAddr)
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
    _, err := io.Copy(w, r)
    errc <- err
}

func chat(a, b io.ReadWriteCloser) {
    fmt.Fprintln(a, "Found one! Say hi.")
    fmt.Fprintln(b, "Found one! Say hi.")

    // It's important to clean up when the conversation is over.
    // To do this we send the error value from each io.Copy call to a channel,
    //  log any non-nil errors, and close both connections. 
    errc := make(chan error, 1)
    go cp(a, b, errc)
    go cp(b, a, errc)
    if err := <-errc; err != nil {
        log.Println(err)
    }
    a.Close()
    b.Close()
}



func match(c io.ReadWriteCloser) {
    fmt.Fprint(c, "Waiting for a partner...")
    select {
    case partner <- c:
        // handled by the other goroutine
    case p := <-partner:
        chat(p, c)
    }
}


func main() {
    http.HandleFunc("/", rootHandler)
    http.Handle("/socket", websocket.Handler(socketHandler))
    err := http.ListenAndServe(listenAddr, nil)
    if err != nil {
        log.Fatal(err)
    }
}

*/