package routes

import (
    "google.golang.org/appengine"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
)

var dev bool
var pages struct {
    MethodNotAllowed []byte
    NotFound []byte
    NotImplemented []byte
}

func init() {
    dev = appengine.IsDevAppServer()
    f, err := ioutil.ReadFile("frontend/method-not-allowed.html")
	if err != nil {
		log.Fatal(err)
	}
    pages.MethodNotAllowed = f
    f, err = ioutil.ReadFile("frontend/not-found.html")
	if err != nil {
		log.Fatal(err)
	}
    pages.NotFound = f
    f, err = ioutil.ReadFile("frontend/not-implemented.html")
	if err != nil {
		log.Fatal(err)
	}
    pages.NotImplemented = f
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    if dev {
        serve(w, r, r.URL)
        return
    }
    url := r.URL
    switch url.Hostname() {
    case "www.bobkidbob.com":
        serve(w, r, url)
    case "bobkidbob.com":
        http.Redirect(w, r, "https://www.bobkidbob.com" + pathAndQuery(url), http.StatusMovedPermanently)
    case "bobkidbob.co.uk":
        http.Redirect(w, r, "https://www.bobkidbob.co.uk" + pathAndQuery(url), http.StatusMovedPermanently)
    case "bobkidbob.info":
        http.Redirect(w, r, "https://www.bobkidbob.info" + pathAndQuery(url), http.StatusMovedPermanently)
    case "bobkidbob.net":
        http.Redirect(w, r, "https://www.bobkidbob.net" + pathAndQuery(url), http.StatusMovedPermanently)
    case "bobkidbob.org":
        http.Redirect(w, r, "https://www.bobkidbob.org" + pathAndQuery(url), http.StatusMovedPermanently)
    default:
        http.Redirect(w, r, "https://www.bobkidbob.com" + pathAndQuery(url), http.StatusFound)
    }
}

func serve(w http.ResponseWriter, r *http.Request, url *url.URL) {
    switch r.Method {
    case http.MethodGet:
        head(w, r, url)
        get(w, url)
    case http.MethodHead:
        head(w, r, url)
    case http.MethodPost:
        methodNotAllowed(w)
    case http.MethodPut:
        methodNotAllowed(w)
    case http.MethodPatch:
        methodNotAllowed(w)
    case http.MethodDelete:
        methodNotAllowed(w)
    case http.MethodConnect:
        methodNotAllowed(w)
    case http.MethodOptions:
        methodNotAllowed(w)
    case http.MethodTrace:
        methodNotAllowed(w)
    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write(pages.NotImplemented)
    }
}

func head(w http.ResponseWriter, r *http.Request, url *url.URL) {
    switch url.Path {
    case "/":
        http.Redirect(w, r, "https://www.linkedin.com/in/bobkidbob/", http.StatusFound)
    default:
        w.WriteHeader(http.StatusNotFound)
    }
}

func get(w http.ResponseWriter, url *url.URL) {
    switch url.Path {
    case "/":
        return
    default:
        w.Write(pages.NotFound)
    }
}

func methodNotAllowed(w http.ResponseWriter) {
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write(pages.MethodNotAllowed)
}

func pathAndQuery(url *url.URL) string {
    s := url.EscapedPath()
    if url.RawQuery != "" {
        s += "?" + url.RawQuery
    }
    return s
}