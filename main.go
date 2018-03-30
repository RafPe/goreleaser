package main

import (
	"github.com/apex/log"
	"github.com/goreleaser/goreleaser/config"
)

func main() {

	config, _ := config.Load("/tmp/goreleaser.yml")

	log.WithFields(log.Fields{
		"Host":  config.Release.Provider.Host,
		"Name":  config.Release.Provider.Name,
		"Owner": config.Release.Provider.Owner,
		"Type":  config.Release.Provider.Type,
	}).Info("Provider")

	// cfg, err := loadConfig(options.Config)
	// if err != nil {
	// 	return err
	// }
}

// func loadConfig(path string) (config.Project, error) {
// 	if path != "" {
// 		return config.Load(path)
// 	}
// 	for _, f := range [4]string{
// 		".goreleaser.yml",
// 		".goreleaser.yaml",
// 		"goreleaser.yml",
// 		"goreleaser.yaml",
// 	} {
// 		proj, err := config.Load(f)
// 		if err != nil && os.IsNotExist(err) {
// 			continue
// 		}
// 		return proj, err
// 	}
// 	// the user didn't specified a config file and the known files
// 	// doest not exist, so, return an empty config and a nil err.
// 	log.Warn("could not load config, using defaults")
// 	return config.Project{}, nil
// }
