package httprouter

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	ROOT = "/"
)

func NewTree() *Tree {
	return &Tree{
		node: &Route{
			Label: ROOT,
			Child: make(map[string]*Route),
		},
	}
}

func (t *Tree) Insert(path string, handler http.Handler, methods ...string) {
	actualRoute := t.node
	if path == ROOT {
		actualRoute.Methods = append(actualRoute.Methods, methods...)
		actualRoute.Handle = handler
	} else {
		roads := strings.Split(path, "/")
		for _, routeName := range roads[1:] {
			NextRoute, ok := actualRoute.Child[routeName]
			if ok {
				actualRoute = NextRoute
			}
			if !ok {
				actualRoute.Child[routeName] = NewRoute(routeName, handler, methods...)
				actualRoute = actualRoute.Child[routeName]
			}
		}
	}
}

func (t *Tree) Search(method string, path string) (http.Handler, error) {
	actualRoute := t.node
	if path != ROOT {
		roads := strings.Split(path, "/")
		for _, routeName := range roads[1:] {
			nextRoute, ok := actualRoute.Child[routeName]
			if !ok {
				if routeName == actualRoute.Label {
					break
				} else {
					err := errors.New(ROUTE_NOT_FOUND)
					return nil, err
				}
			}
			actualRoute = nextRoute
		}
	}
	err := actualRoute.IsAllowed(method)
	if err != nil {
		
		return nil, err
	}
	fmt.Println(actualRoute.Label)
	return actualRoute.Handle, nil
}
