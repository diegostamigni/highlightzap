[![Go Report](https://goreportcard.com/badge/github.com/diegostamigni/highlightzap)](https://goreportcard.com/report/github.com/diegostamigni/highlightzap)

# Highlight Zap

A simple zapcore.Core implementation to integrate with Highlight.

To use, initialize highlight like normal, create a new HighlightCore, then wrap with a NewTee. [See the example code](example/main.go) for a detailed example.

## Testing 

To test this code use `HL_PROJECT_ID=MY_HIGHLIGHT_TOKEN go test`
