package types

// The program's config (to parse into).
type Config struct {
	Delay   float64 `json:"delay"`
	Dir     string  `json:"dir"`
	SVCFile string  `json:"service_file_dir"`
}
