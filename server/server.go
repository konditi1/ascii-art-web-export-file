package webart

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	webart "webart/ascii"
)

var (
	filename    string
	fileContent string
)

type Result struct {
	Result string
}

type Error_template struct {
	Title   string
	Message string
	Status  int
}

var Data Result

// ProcessForm handles the data passed to the form
func ProcessForm(w http.ResponseWriter, r *http.Request, filename string) {
	Data = Result{}
	var art string
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		ShowError(w, "server error", " Page Not Found", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		ShowError(w, "server error", "Internal Server Error", http.StatusInternalServerError)
		return
	}
	text := r.Form.Get("input")
	banner := r.Form.Get("banner")
	Bannercontent, err := webart.FileReader(banner)
	if err != nil {
		ShowError(w, "server error", "Resource not found or damaged", http.StatusInternalServerError)
		return
	}
	if !webart.CheckAscii(text) {
		ShowError(w, "Bad request", "Bad request", http.StatusBadRequest)
		return
	}
	art, err = webart.Ascii(Bannercontent, text)
	if err != nil {
		ShowError(w, "server error", "Resource not found or damaged", http.StatusInternalServerError)
		return
	}
	Data = Result{Result: art}

	fileContent = Data.Result

	if r.Form.Get("download") != "" {
    // Redirect the user to the `/save` path to trigger the download
    http.Redirect(w, r, "/save", http.StatusSeeOther)
    return
}
	w.WriteHeader(http.StatusOK)
	t.Execute(w, Data)
}

// ShowError handles all errors encountered
func ShowError(writer http.ResponseWriter, title string, message string, code int) {
	tmpl, _ := template.ParseFiles("template/error.html")
	Error_case := Error_template{Title: title, Message: message, Status: code}
	writer.WriteHeader(code)
	tmpl.Execute(writer, Error_case)
}

func Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ShowError(w, "Method Not Allowed", "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	timeStamp := time.Now().Format("2006012150405")
	filename = "asciiArt-" + timeStamp + ".txt"

	w.Header().Set("content-Disposition", "attachment; filename="+filename)
	contentLenth := len(fileContent)
	w.Header().Set("Content-Length", strconv.Itoa(contentLenth))
	w.Header().Set("content-Type", "text/plain")

	_, err := w.Write([]byte(Data.Result))
	if err != nil {
		ShowError(w, "server error", "Internal server Error", http.StatusInternalServerError)
		return
	}
}
