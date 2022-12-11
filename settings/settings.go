package settings

import "flag"

type Settings struct {
	Port uint
	Path string
}

func Parse() Settings {
	var (
		port = flag.Uint("port", 3000, "Server port")
		path = flag.String("path", ".", "Path to a folder")
	)

	flag.Parse()

	return Settings{
		Port: *port,
		Path: *path,
	}
}
