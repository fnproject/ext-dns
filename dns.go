package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

func init() {
	server.RegisterExtension(&Dns{})
}

const (
	EnvAPIHost = "API_HOST"
)

type Dns struct {
}

func (e *Dns) Name() string {
	return "github.com/treeder/fn-ext-dns"
}

func (e *Dns) Setup(s fnext.ExtServer) error {
	fmt.Println("SETTING UP DNS")
	if os.Getenv(EnvAPIHost) == "" {
		return fmt.Errorf("%s env var is required for dns extension", EnvAPIHost)
	}
	s.AddRootMiddleware(&Middleware{})
	return nil
}

type Middleware struct {
}

func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// log := ctx.Value("logger").(logrus.FieldLogger)
		fmt.Println("DNS middleware called")

		// ensure API_HOST env var is set
		apiHost := os.Getenv("API_HOST")
		if apiHost == "" {
			panic(errors.New("API_HOST environment variable is not. It is required for the DNS middleware."))
		}

		host := r.Host
		apiHosts := strings.Split(apiHost, ",")
		for _, h := range apiHosts {
			if strings.HasPrefix(strings.TrimSpace(h), host) { // HasPrefix to ignore port
				// do the regular thing
				next.ServeHTTP(w, r)
				return
			}
		}

		// pull appname out of URL, currently will only match the first part of the domain name
		// eg: myapp.domain.com will get myapp
		split := strings.Split(host, ".")
		fmt.Println("split:", split)
		if len(split) < 3 {
			// just pass it along, could be localhost or something else
			next.ServeHTTP(w, r)
			return
		}
		appName := split[0]

		// TODO: ASAP: make these keys part of fn package context keys, eg: fn.AppNameKey, fn.Path
		ctx = context.WithValue(ctx, "app_name", appName)
		// would be nice to not have to set path here, the function call handler can just get it from the request
		ctx = context.WithValue(ctx, "path", r.URL.Path)
		mctx := fnext.GetMiddlewareController(ctx)
		mctx.CallFunction(w, r.WithContext(ctx))
	})
}
