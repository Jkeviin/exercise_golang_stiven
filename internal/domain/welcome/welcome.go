package welcome

type Welcome struct {
	Message   string   `json:"CARRO"`
	Version   float32  `json:"version"`
	Endpoints []string `json:"endpoints"`
}
