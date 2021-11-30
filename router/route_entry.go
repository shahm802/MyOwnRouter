package router

import (
	"net/http"
	"regexp"
)

type RouteEntry struct {
	Path    *regexp.Regexp
	Method  string
	Handler http.HandlerFunc
}

func (re *RouteEntry) Match(r *http.Request) map[string]string {
	match := re.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil
	}

	params := make(map[string]string)
	groupNames := re.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}
	return params

}
