package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"path"
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
	checkHost := r.FormValue("host")

	commandString := fmt.Sprintf("/usr/bin/curl -ks %s -v", checkHost)

	type replyResponse struct {
		Resp string
		Out  string
	}
	thisOut := replyResponse{}

	cmd := exec.Command("/bin/bash", "-c", commandString)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		thisOut.Resp = "dead"
	} else {
		thisOut.Resp = "alive"
		thisOut.Out = string(stdout)
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
