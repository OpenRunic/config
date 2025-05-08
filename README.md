![Test Status](https://github.com/OpenRunic/config/actions/workflows/master-push.yml/badge.svg)

## Configurator Library for GO

This library allow you to load configurations from flags, env, json, yaml and much more through custom interface implementation.

```
import (
  ...

	"github.com/OpenRunic/config"
	"github.com/OpenRunic/config/options"
	"github.com/OpenRunic/config/env"
	"github.com/OpenRunic/config/json"
)

type ServerConfig struct {
  Host string
  Port int
}

var sc ServerConfig
_, err := config.Parse(
  options.New(),
  &sc,
  config.Register(json.New()),
  config.Register(env.New()),
  config.Add("host", "localhost", "Server hostname"),
  config.Add("port", 3000, "Server port"),
)

if err != nil {
  log.Fatal(err)
}

println("Server Host: ", sc.Host)
println("Server Port: ", sc.Port)
```
