package main

import (
	"encoding/json"
	"github.com/dnguyenzd/task-schema-test/internal/task"
	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
	"net/http"
)

func main() {
	renderer := render.New()
	r := chi.NewRouter()
	r.Get("/raw-message", func(w http.ResponseWriter, r *http.Request) {
		data := task.CustomResourceSpec{
			Name: "custom-resource",
		}
		marshalled, err := json.Marshal(data)
		if err != nil {
			renderer.Text(w, http.StatusInternalServerError, err.Error())
			return
		}

		renderer.JSON(w, http.StatusOK, task.WithRawMessage{
			Data: marshalled,
		})
	})

	r.Get("/any", func(w http.ResponseWriter, r *http.Request) {
		renderer.JSON(w, http.StatusOK, task.WithAny{
			Data: task.CustomResourceSpec{
				Name: "custom-resource",
			},
		})
	})

	http.ListenAndServe(":3000", r)
}
