package router

import "encoding/json"

type (
	Body struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
		Metadata   interface{} `json:"metadata,omitempty"`
	}
)

func (b *Body) ToJson() string {
	bytes, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
