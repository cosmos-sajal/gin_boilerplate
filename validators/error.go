package validators

type Error struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type RestError struct {
	Status int     `json:"status"`
	Errors []Error `json:"errors"`
}
