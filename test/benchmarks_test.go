package test

import (
	"html/template"
	"net/http"
	"os"
	"testing"

	"github.com/go-ego/ego"
)

func BenchmarkOneRoute(B *testing.B) {
	router := ego.New()
	router.GET("/ping", func(c *ego.Context) {})
	runRequest(B, router, "GET", "/ping")
}

func BenchmarkRecoveryMiddleware(B *testing.B) {
	router := ego.New()
	router.Use(ego.Recovery())
	router.GET("/", func(c *ego.Context) {})
	runRequest(B, router, "GET", "/")
}

func BenchmarkLoggerMiddleware(B *testing.B) {
	router := ego.New()
	router.Use(ego.LoggerWithWriter(newMockWriter()))
	router.GET("/", func(c *ego.Context) {})
	runRequest(B, router, "GET", "/")
}

func BenchmarkManyHandlers(B *testing.B) {
	router := ego.New()
	router.Use(ego.Recovery(), ego.LoggerWithWriter(newMockWriter()))
	router.Use(func(c *ego.Context) {})
	router.Use(func(c *ego.Context) {})
	router.GET("/ping", func(c *ego.Context) {})
	runRequest(B, router, "GET", "/ping")
}

func Benchmark5Params(B *testing.B) {
	ego.DefaultWriter = os.Stdout
	router := ego.New()
	router.Use(func(c *ego.Context) {})
	router.GET("/param/:param1/:params2/:param3/:param4/:param5", func(c *ego.Context) {})
	runRequest(B, router, "GET", "/param/path/to/parameter/john/12345")
}

func BenchmarkOneRouteJSON(B *testing.B) {
	router := ego.New()
	data := struct {
		Status string `json:"status"`
	}{"ok"}
	router.GET("/json", func(c *ego.Context) {
		c.JSON(200, data)
	})
	runRequest(B, router, "GET", "/json")
}

var htmlContentType = []string{"text/html; charset=utf-8"}

func BenchmarkOneRouteHTML(B *testing.B) {
	router := ego.New()
	t := template.Must(template.New("index").Parse(`
		<html><body><h1>{{.}}</h1></body></html>`))
	router.SetHTMLTemplate(t)

	router.GET("/html", func(c *ego.Context) {
		c.HTML(200, "index", "hola")
	})
	runRequest(B, router, "GET", "/html")
}

func BenchmarkOneRouteSet(B *testing.B) {
	router := ego.New()
	router.GET("/ping", func(c *ego.Context) {
		c.Set("key", "value")
	})
	runRequest(B, router, "GET", "/ping")
}

func BenchmarkOneRouteString(B *testing.B) {
	router := ego.New()
	router.GET("/text", func(c *ego.Context) {
		c.String(200, "this is a plain text")
	})
	runRequest(B, router, "GET", "/text")
}

func BenchmarkManyRoutesFist(B *testing.B) {
	router := ego.New()
	router.Any("/ping", func(c *ego.Context) {})
	runRequest(B, router, "GET", "/ping")
}

func BenchmarkManyRoutesLast(B *testing.B) {
	router := ego.New()
	router.Any("/ping", func(c *ego.Context) {})
	runRequest(B, router, "OPTIONS", "/ping")
}

func Benchmark404(B *testing.B) {
	router := ego.New()
	router.Any("/something", func(c *ego.Context) {})
	router.NoRoute(func(c *ego.Context) {})
	runRequest(B, router, "GET", "/ping")
}

func Benchmark404Many(B *testing.B) {
	router := ego.New()
	router.GET("/", func(c *ego.Context) {})
	router.GET("/path/to/something", func(c *ego.Context) {})
	router.GET("/post/:id", func(c *ego.Context) {})
	router.GET("/view/:id", func(c *ego.Context) {})
	router.GET("/favicon.ico", func(c *ego.Context) {})
	router.GET("/robots.txt", func(c *ego.Context) {})
	router.GET("/delete/:id", func(c *ego.Context) {})
	router.GET("/user/:id/:mode", func(c *ego.Context) {})

	router.NoRoute(func(c *ego.Context) {})
	runRequest(B, router, "GET", "/viewfake")
}

type mockWriter struct {
	headers http.Header
}

func newMockWriter() *mockWriter {
	return &mockWriter{
		http.Header{},
	}
}

func (m *mockWriter) Header() (h http.Header) {
	return m.headers
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockWriter) WriteHeader(int) {}

func runRequest(B *testing.B, r *ego.Engine, method, path string) {
	// create fake request
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	w := newMockWriter()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		r.ServeHTTP(w, req)
	}
}
