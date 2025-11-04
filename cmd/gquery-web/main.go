package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>HTTP Request Information</title>
	<style>
		table { margin-bottom: 20px; border-collapse: collapse; }
		th { background-color: #f0f0f0; text-align: left; padding: 8px; }
		td { padding: 8px; }
	</style>
</head>
<body>
	<h1>HTTP Request Information</h1>

	<h2>HTTP Headers</h2>
	<table border="1">
		<tr>
			<th>Header</th>
			<th>Values</th>
		</tr>
		{{range $key, $values := .Headers}}
		<tr>
			<td>{{$key}}</td>
			<td>{{range $values}}{{.}} {{end}}</td>
		</tr>
		{{end}}
	</table>

	<h2>Query Parameters</h2>
	<table border="1">
		<tr>
			<th>Parameter</th>
			<th>Values</th>
		</tr>
		{{range $key, $values := .Params}}
		<tr>
			<td>{{$key}}</td>
			<td>{{range $values}}{{.}} {{end}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("params").Parse(htmlTemplate))
	r.ParseForm()

	data := struct {
		Headers map[string][]string
		Params  map[string][]string
	}{
		Headers: r.Header,
		Params:  r.Form,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on http://localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
