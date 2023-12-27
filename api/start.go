package api

import (
	"context"
	"encoding/json"
	"net/http"
	"radim91/entity"
	"regexp"
	"strings"
)

type ctxKey struct{}
type route struct {
    method  string
    regex   *regexp.Regexp
    handler http.HandlerFunc
}

var routes = []route{
    newRoute("GET", "projects", getProjects),
    newRoute("GET", "containers", getContainers),
    newRoute("GET", "images", getImages),
    newRoute("GET", "networks", getNetworks),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
    return route{
        method: method,
        regex: regexp.MustCompile(pattern),
        handler: handler,
    }
}


func getField(r *http.Request, index int) string {
    fields := r.Context().Value(ctxKey{}).([]string)
    return fields[index]
}

func Serve(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")

    var allow []string
    for _, route := range routes {
        matches := route.regex.FindStringSubmatch(r.URL.Path)
        if len(matches) > 0 {
            if r.Method != route.method {
                allow = append(allow, route.method)
                continue
            }
            
            ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
            route.handler(w, r.WithContext(ctx))
            return
        }
    }

    if len(allow) > 0 {
        w.Header().Set("Allow", strings.Join(allow, ", "))
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        return
    }

    http.NotFound(w, r)
}

func getProjects(w http.ResponseWriter, r *http.Request) {
    projects := entity.GetProjects()
    jData, err := json.Marshal(projects)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)
}

func getContainers(w http.ResponseWriter, r *http.Request) {
    containers := entity.GetContainers()
    jData, err := json.Marshal(containers)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)
}

func getImages(w http.ResponseWriter, r *http.Request) {
    images := entity.GetImages()
    jData, err := json.Marshal(images)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)
}

func getNetworks(w http.ResponseWriter, r *http.Request) {
    networks := entity.GetNetworks()
    jData, err := json.Marshal(networks)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jData)
}
