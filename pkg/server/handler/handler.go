package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/server/httpmethod"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

// Initialize setup handler
func Initialize(schema graphql.Schema, method httpmethod.HTTPMethod) error {
	switch method {
	case httpmethod.GET:
		http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			result := executeQuery(r.URL.Query().Get("query"), schema)
			json.NewEncoder(w).Encode(result)
		})
	case httpmethod.POST:
		http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
			var p postData
			if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
				w.WriteHeader(400)
				return
			}
			result := graphql.Do(graphql.Params{
				Context:        req.Context(),
				Schema:         schema,
				RequestString:  p.Query,
				VariableValues: p.Variables,
				OperationName:  p.Operation,
			})
			if err := json.NewEncoder(w).Encode(result); err != nil {
				fmt.Printf("could not write result to response: %s", err)
			}
		})
	default:
		return errors.Errorf("invalid http method: %s", method)
	}
	return nil
}
