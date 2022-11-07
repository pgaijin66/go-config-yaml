## Reading configuration from yaml file

#### Usage

```
go run cmd/main.go
```

### Note

When using `viper.AutomaticEnv()`, it checks for env variables of the same key and if its defined the overrides it with the value from env variable. `viper.AutomaticEnv()` is case sensitive. it checks for environment varible but uppercased verison i.e if you want to override `environment` in config file you need to set `export ENVIRONMENT="PROD"` in your shell.

## Gotcha

For nested yaml file, viper needs to replace `.` with `_` and you need to set environment variable like this

For example

if you configuration is like this
```
server:
  host: "localhost"
  port: 9090
```

viper looks for `SERVER.PORT` but since shell does not support `dot` format for environment variables, we transform it to `dash` format using 

`
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
`

and if you want to override port, then your environment variable should be `export SERVER_PORT=9999"`
