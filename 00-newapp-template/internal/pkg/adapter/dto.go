package adapter

type Gopher struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Things      map[string]Thing `json:"things"`
}

type Thing struct {
	Gopher      Gopher `json:"gopher"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
