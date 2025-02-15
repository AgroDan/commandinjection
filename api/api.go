package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"path"
	"strings"
)

// This package will host the API that will run as a webserver
// inside a docker container. Ultimately, it will contain a "welcome"
// landing page, with a link to another form that will ping something.
// The ping will be a binary in the container that will be called from
// essentially a shell environment.

func Run(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", showIndex)
	mux.HandleFunc("/hostalive", hostAliveFE)
	mux.HandleFunc("/checkhost", hostAlive)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func showIndex(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hostAliveFE(w http.ResponseWriter, r *http.Request) {
	// This is the front-end to the ping command. I'll worry about
	// combining this func with the above later because this is code
	// duplication.
	fp := path.Join("templates", "hostalive.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hostAlive(w http.ResponseWriter, r *http.Request) {
	pingHost := r.FormValue("host")
	cmd := exec.Command(fmt.Sprintf("ping -c 1 %s", pingHost))
	stdout, err := cmd.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type replyResponse struct {
		Resp string
	}

	strOut := string(stdout)
	thisOut := replyResponse{}
	if strings.Contains(strOut, "1 packets transmitted, 1 received") {
		thisOut.Resp = "alive"
	} else {
		thisOut.Resp = "dead"
	}

	fp := path.Join("templates", "response.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, thisOut); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
