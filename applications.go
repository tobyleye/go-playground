package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"log"
	"sync"
)


func cmdArgs() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + strings.ToUpper(arg)
		sep = " "
	}
	fmt.Println(s)
}

func tellStory() {

	fmt.Println("Tell us a story. \nExit by pressing :q on a new line.")
	f := bufio.NewScanner(os.Stdin)

	var story string

	for f.Scan() {
		line := f.Text()
		if line == ":q" {
			break
		}

		story += line + "\n"
	}

	fmt.Printf("Thank you for sharing :D")
}

func dup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if  n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} 
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		fmt.Printf("data:", string(data))
		if err != nil {
			fmt.Fprint(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func normalizeUrl(url string) string {
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	return url
}


func fetch() {
	for _, url := range os.Args[1:] {
		url = normalizeUrl(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("Status code: %v\n", resp.Status)
		bytes, err := io.Copy(os.Stdout, resp.Body)
		// resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("number of bytes written is %d", bytes)
	}
}


func singlefetch(url string, ch chan<- string) {
	start := time.Now()
	url = normalizeUrl(url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}


func fetchall() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go singlefetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}


func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()	
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	
				<div>
					<h1>Hello world</h1>
					<p>In case you are wondering, my name is Tobi, and i'm a fullstack software engineer
						with some experience in typescript, react and now i'm learning Go.
					</p>
				</div>
			
	`)
}


var mu sync.Mutex
var count int 

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}

func showStats(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k,v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k,v)
	}
}


func webServer() {
	http.HandleFunc("/me", aboutMe)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/stats", showStats)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}



func main() {
	// cmdArgs()
	// dup()
	// tellStory()
	// dup2()
	// dup3()
	// fetchall()
	webServer()
}


