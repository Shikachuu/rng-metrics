# RNG Metrics

This little utility container generates random log messages with go's new `slog` library and prometheus metrics about them.

In the future maybe it will provide some random traces or the ability to give a random dataset and whatnot, to give users the felxiblity to test their (LGTM)[https://github.com/grafana#lgtm] or any other monitroing infrastructure with some periodically generated test data.

## Available flags

You can probably figure it out from the --help command, but here we go:
| Flag | Short-hand | What it does |
| --- | --- | --- |
| `--json` | none | changes the logging fromat from logfmt to json, default is `false` |
| `--duration` | `-d` | changes the generation period only accepts a valid go duration (eg. `10s`), default 15 sec |
| `--addr` | `-a` | changes the address the `/metrics` endpoint listening on, default `:8080` |
