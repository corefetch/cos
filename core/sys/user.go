package sys

type User struct {
	ID    int64             `json:"id"`
	Login string            `json:"login"`
	Names []string          `json:"names"`
	Meta  map[string]string `json:"meta"`
}
