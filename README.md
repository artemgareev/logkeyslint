# logkeyslint

logkeyslint is a go static analysis tool to find rs/zerolog log keys types mismatches

Example of log key mismatch:
```go
    log.Info().Str("device_id", "1").Msg("device_id passed as string")
    log.Info().Int("device_id", 1).Msg("device_id passed as int")
```


To use this:
### Create the Plugin From This Linter

1. Download the source code \*
2. From the root project directory, run `go build -buildmode=plugin plugin/example.go`.
3. Copy the generated `example.so` file into your project or to some other known location of your choosing. \**


### Create a Copy of `golangci-lint` that Can Run with Plugins

In order to use plugins, you'll need a golangci-lint executable that can run them. The normal version of this project 
is built with the vendors option, which breaks plugins that have overlapping dependencies.

1. Download [golangci-lint](https://github.com/golangci/golangci-lint) source code
2. From the projects root directory, run `make vendor_free_build`
3. Copy the `golangci-lint` executable that was created to your path, project, or other location

### Configure Your Project for Linting

If you already have a linter plugin available, you can follow these steps to define it's usage in a projects 
`.golangci.yml` file. An example linter can be found at [here](https://github.com/golangci/example-plugin-linter). If you're looking for 
instructions on how to configure your own custom linter, they can be found further down.

1. If the project you want to lint does not have one already, copy the [.golangci.yml](https://github.com/golangci/golangci-lint/blob/master/.golangci.yml) to the root directory.
2. Adjust the yaml to appropriate `linters-settings:custom` entries as so:
```
linters-settings:
 custom:
  example:
   path: /example.so
   description: The description of the linter
   original-url: github.com/golangci/example-linter
```

That is all the configuration that is required to run a custom linter in your project. Custom linters are enabled by default,
but abide by the same rules as other linters. If the disable all option is specified either on command line or in 
`.golang.yml` files `linters:disable-all: true`, custom linters will be disabled; they can be re-enabled by adding them 
to the `linters:enable` list, or providing the enabled option on the command line, `golangci-lint run -Eexample`.
