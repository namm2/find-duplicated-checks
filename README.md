# Find duplicated Pingdom checks

This tool get the list of uptime checks and check for duplicated items.

## Prerequisites

Generate Pingdom API token at: https://my.pingdom.com/app/api-tokens

Provide it in runtime:
```
PINGDOM_API_TOKEN=loremipsum ./checks
```

## Build

```
go build .
```
