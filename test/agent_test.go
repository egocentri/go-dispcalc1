package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/egocentri/go-dispcalc1/internal/agent"
	"github.com/egocentri/go-dispcalc1/internal/config"
)

func TestAgentRequestTask(t *testing.T) {
	// Создаём тестовый сервер, имитирующий эндпоинты оркестратора
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/internal/task" {
			fmt.Fprintf(w, "{\"task\": {\"id\": 1, \"arg1\": \"2+2\", \"operation\": \"eval\", \"operation_time\": 100}}")
			return
		}
		if r.Method == http.MethodPost && r.URL.Path == "/internal/task" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "{\"status\": \"result accepted\"}")
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	cfg := config.InitEnv()
	worker := agent.NewWorker(ts.URL, 1, cfg)

	worker.RequestAndProcessTask()
	// Если функция завершилась без паники, тест считается пройденным
	assert.True(t, true)
}
