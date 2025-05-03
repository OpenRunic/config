![Test Status](https://github.com/OpenRunic/config/actions/workflows/master-push.yml/badge.svg)

## Configurator Library for GO

This library allow you to load configurations from flags, environment, json file and much more through custom interface implementation.

```
type ConfigData struct {
  Host string
  Port int
}

var configs ConfigData
_, err := config.Default(&configs,
  config.Add("host", "localhost", "Server hostname"),
  config.Add("port", 3000, "Server port"),
)

if err != nil {
  log.Fatal(err)
}

println("Server Host: ", configs.Host)
println("Server Port: ", configs.Port)
```
