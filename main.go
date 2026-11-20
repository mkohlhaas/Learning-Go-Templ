package main

import (
	"log"
	"net/http"
	"sync"

	atempl "github.com/a-h/templ"
	"go-templ/templ"
)

// clickCounter is an in-memory counter used by lesson 8 to demonstrate
// templ + HTMX swapping fragments without a page reload.
var (
	clickMu   sync.Mutex
	clickCnt  int
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Course index.
	mux.HandleFunc("/", homeHandler)

	// Lessons.
	mux.HandleFunc("/lesson1", lesson1Handler)
	mux.HandleFunc("/lesson2", lesson2Handler)
	mux.HandleFunc("/lesson3", lesson3Handler)
	mux.HandleFunc("/lesson4", lesson4Handler)
	mux.HandleFunc("/lesson5", lesson5Handler)
	mux.HandleFunc("/lesson6", lesson6Handler)
	mux.HandleFunc("/lesson7", lesson7Handler)
	mux.HandleFunc("/lesson8", lesson8Handler)

	// HTMX fragment endpoint for lesson 8.
	mux.HandleFunc("/lesson8/click", lesson8ClickHandler)

	log.Println("Listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// render wraps a component in the shared Layout and writes it to the response.
// The component is supplied to Layout as context children so the shared shell
// can render it with { children... }.
func render(w http.ResponseWriter, r *http.Request, title string, children atempl.Component, useHTMX bool) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx := atempl.WithChildren(r.Context(), children)
	if err := templ.Layout(title, useHTMX).Render(ctx, w); err != nil {
		log.Printf("templ render error: %v", err)
		http.Error(w, "render error", http.StatusInternalServerError)
	}
}

// ---------------------------------------------------------------------------
// Lesson handlers
// ---------------------------------------------------------------------------

func homeHandler(w http.ResponseWriter, r *http.Request) {
	lessons := []templ.Lesson{
		{Path: "/lesson1", Title: "Lesson 1: First component", Desc: "A Go function that returns HTML — no template files."},
		{Path: "/lesson2", Title: "Lesson 2: Props", Desc: "Pass data into components as Go function parameters."},
		{Path: "/lesson3", Title: "Lesson 3: Conditionals", Desc: "Show or hide content with @if / @else."},
		{Path: "/lesson4", Title: "Lesson 4: Loops", Desc: "Render lists with @for over a Go slice."},
		{Path: "/lesson5", Title: "Lesson 5: Composition", Desc: "Nest components and pass components as children."},
		{Path: "/lesson6", Title: "Lesson 6: Serving over HTTP", Desc: "Render a component from a net/http handler."},
		{Path: "/lesson7", Title: "Lesson 7: Forms", Desc: "POST a form and re-render with submitted data."},
		{Path: "/lesson8", Title: "Lesson 8: templ + HTMX", Desc: "Return HTML fragments that HTMX swaps into the page."},
	}
	render(w, r, "Home", templ.Home(lessons), false)
}

func lesson1Handler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "Lesson 1: First component", templ.Lesson1(), false)
}

func lesson2Handler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "Lesson 2: Props", templ.Lesson2("Ada", 3), false)
}

func lesson3Handler(w http.ResponseWriter, r *http.Request) {
	// Toggle with ?loggedin=1 to see both branches.
	loggedIn := r.URL.Query().Get("loggedin") == "1"
	render(w, r, "Lesson 3: Conditionals", templ.Lesson3(loggedIn, "Pro"), false)
}

func lesson4Handler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "Lesson 4: Loops", templ.Lesson4(), false)
}

func lesson5Handler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "Lesson 5: Composition", templ.Lesson5(), false)
}

func lesson6Handler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "Lesson 6: Serving over HTTP", templ.Lesson6("Hello from a net/http handler."), false)
}

func lesson7Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		render(w, r, "Lesson 7: Forms", templ.Lesson7(name, true), false)
		return
	}
	render(w, r, "Lesson 7: Forms", templ.Lesson7("", false), false)
}

func lesson8Handler(w http.ResponseWriter, r *http.Request) {
	clickMu.Lock()
	cnt := clickCnt
	clickMu.Unlock()
	render(w, r, "Lesson 8: templ + HTMX", templ.Lesson8(cnt), true)
}

func lesson8ClickHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	clickMu.Lock()
	clickCnt++
	cnt := clickCnt
	clickMu.Unlock()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templ.CounterFragment(cnt).Render(r.Context(), w); err != nil {
		log.Printf("templ render error: %v", err)
		http.Error(w, "render error", http.StatusInternalServerError)
	}
}
