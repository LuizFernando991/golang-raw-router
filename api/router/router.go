package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	regex      *regexp.Regexp
	paramNames []string
}

type Router struct {
	routes []Route
}

type RouteGroup struct {
	prefix      string
	middlewares []func(http.HandlerFunc) http.HandlerFunc
	router      *Router
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	regexPattern := "^" + regexp.MustCompile(`\{(\w+)\}`).ReplaceAllStringFunc(pattern, func(m string) string {
		param := m[1 : len(m)-1]
		return `(?P<` + param + `>[^/]+)`
	}) + "$"

	regex := regexp.MustCompile(regexPattern)

	paramNames := []string{}
	matches := regexp.MustCompile(`\{(\w+)\}`).FindAllStringSubmatch(pattern, -1)
	for _, match := range matches {
		paramNames = append(paramNames, match[1])
	}

	r.routes = append(r.routes, Route{
		Method:     method,
		Pattern:    pattern,
		Handler:    handler,
		regex:      regex,
		paramNames: paramNames,
	})
}

func (r *Router) Group(prefix string, middlewares ...func(http.HandlerFunc) http.HandlerFunc) *RouteGroup {
	return &RouteGroup{
		prefix:      prefix,
		middlewares: middlewares,
		router:      r,
	}
}

func (g *RouteGroup) Group(prefix string, middlewares ...func(http.HandlerFunc) http.HandlerFunc) *RouteGroup {
	return &RouteGroup{
		prefix:      g.prefix + prefix,
		middlewares: append(g.middlewares, middlewares...),
		router:      g.router,
	}
}

func (g *RouteGroup) Handle(method, pattern string, handler http.HandlerFunc, routeMiddlewares ...func(http.HandlerFunc) http.HandlerFunc) {
	fullPattern := g.prefix + pattern

	// apply routes middlewares
	for i := len(routeMiddlewares) - 1; i >= 0; i-- {
		handler = routeMiddlewares[i](handler)
	}

	// apply group middlewares
	for i := len(g.middlewares) - 1; i >= 0; i-- {
		handler = g.middlewares[i](handler)
	}

	g.router.Handle(method, fullPattern, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method != req.Method {
			continue
		}
		if !route.regex.MatchString(req.URL.Path) {
			continue
		}

		match := route.regex.FindStringSubmatch(req.URL.Path)
		params := map[string]string{}
		for i, name := range route.regex.SubexpNames() {
			if i > 0 && name != "" {
				params[name] = match[i]
			}
		}

		ctx := req.Context()
		for key, val := range params {
			ctx = AddURLParam(ctx, key, val)
		}
		req = req.WithContext(ctx)

		route.Handler(w, req)
		return
	}
	http.NotFound(w, req)
}
