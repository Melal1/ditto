package config

import "flag"

func LoadWithFlags() Config {
	lock := flag.Bool("lock", false, "lock settings toggles")
	unlock := flag.Bool("unlock", false, "unlock settings toggles")
	flag.Parse()

	cfg := LoadConfig()
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "lock":
			cfg.Locked = *lock
			SaveConfig(cfg)
		case "unlock":
			cfg.Locked = !*unlock
			SaveConfig(cfg)
		}
	})
	return cfg
}
