package main

import (
	"io"
	"text/template"
)

const indexTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<link rel="stylesheet" href="https://unpkg.com/tachyons@4.10.0/css/tachyons.min.css"/>
    </head>
	<body class="bg-washed-yellow pa4">
	<ul>
	{{- range .}}<li><a href="/tc/{{.User}}">{{.User}}</a></li>{{end -}}
	</ul>
	<div class="bg-washed-green bw3 ma3 center measure">
		<h5>New User</h5>
		<form action="/" method="post">
 			User:<br>
  			<input type="text" name="user"><br>
  			Password:<br>
  			<input type="text" name="password"><br>
			<input type="submit" value="Submit">
		</form> 
	</div>
	</body>
</html>
`

const tcTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<link rel="stylesheet" href="https://unpkg.com/tachyons@4.10.0/css/tachyons.min.css"/>
    </head>
	<body class="bg-washed-yellow pa4">
		<div class="bg-washed-green bw3 ma3 center measure">
			<h5>Punch in time</h5>
			<form action="/tc/{{.User}}" method="post">
 				Week:<br>
  				<input type="text" name="week"><br>
  				Day:<br>
  				<input type="text" name="day"><br>
				Time:<br>
  				<input type="text" name="time"><br>
				Password:<br>
  				<input type="text" name="password"><br>
				<input type="submit" value="Submit">
			</form> 
		</div>
		<h1>{{.User}}</h1>
		<!-- In case you forget... {{.Password}} -->
		<table style="width:100%">
		<tr>
			<th>Sunday</th>
			<th>Monday</th>
			<th>Tuesday</th>
			<th>Wednesday</th>
			<th>Thursday</th>
			<th>Friday</th>
			<th>Saturday</th>
		</tr>
		{{- range .Time -}}
			<tr>
			{{range .}}<th>{{.}}</th>{{end}}
			</tr>
		{{- end -}}
		</table>
	</body>
</html>
`

func parseIndex(w io.Writer, tcs []*TimeCard) {
	template.Must(template.New("index").Parse(indexTemplate)).Execute(w, tcs)
}

func parseAll(w io.Writer) {
	parseIndex(w, readAll())
}

func parseTC(w io.Writer, tc *TimeCard) {
	template.Must(template.New("tc").Parse(tcTemplate)).Execute(w, tc)
}

func parseDB(w io.Writer, user string) {
	parseTC(w, readDB(user))
}