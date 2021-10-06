# Find duplicated Pingdom checks

This tool get the list of uptime checks and check for duplicated items.

## Prerequisites

Generate Pingdom API token at: https://my.pingdom.com/app/api-tokens

## Build

```
go build .
```

## Usage

Provide Pingdom Token:
```bash
export PINGDOM_API_TOKEN=loremipsum

./find-duplicated-checks
```

To remove all the duplicates:

```bash
export REMOVE_DUPLICATED=True

./find-duplicated-checks
```
