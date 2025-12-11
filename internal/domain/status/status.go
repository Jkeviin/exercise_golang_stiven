package status

type Status struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Uptime  int64  `json:"uptime"`
}

