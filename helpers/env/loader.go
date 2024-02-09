package env

import (
	"simplink/config"

	"github.com/joho/godotenv"
)

func Load(file *string, i ...interface{}) error {
	if file != nil {
		err := godotenv.Load(*file)
		if err != nil {
			return err
		}
	}

	for _, its := range i {
		switch cfg := its.(type) {
		case *config.Mysql:
			if err := cfg.LoadFromEnv(); err != nil {
				return err
			}
		}
	}

	return nil
}
