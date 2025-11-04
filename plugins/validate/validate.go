package traefik_plugin_validate

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Config struct {
	SecretField string `json:"secretField,omitempty"`
	SecretValue string `json:"secretValue,omitempty"`
}

func CreateConfig() *Config {
	return &Config{}
}

type Validate struct {
	next        http.Handler
	secretField string
	secretValue string
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return &Validate{
		next:        next,
		secretField: config.SecretField,
		secretValue: config.SecretValue,
	}, nil
}

func (v *Validate) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var data map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	val, ok := data[v.secretField].(string)
	if !ok || val != v.secretValue {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Invalid secret"))
		return
	}

	v.next.ServeHTTP(rw, req)
}
