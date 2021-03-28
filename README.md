# logkeyslint

logkeyslint is a go static analysis tool to find rs/zerolog log keys types mismatches

Example of log key mismatch:
```go
    log.Info().Str("device_id", "1").Msg("device_id passed as string")
    log.Info().Int("device_id", 1).Msg("device_id passed as int")
```
