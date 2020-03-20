# go-website-lookup-tool
Simple website lookup tool, built with golang

## What can it do
- Lookup for Name Server
- Resolve IP Address
- Lookup for Canonical Name
- Lookup for MX Records

# Notes
- this tool is not completely my own idea, it is created as part of my journey learning golang, and made based on [TutorialEdge](https://www.youtube.com/channel/UCwFl9Y49sWChrddQTD9QhRA) [video](https://www.youtube.com/watch?v=i2p0Snwk4gc)
- this tool is also created as part of my journey learning Computer Network

## Building the tool
```
go build cmd/cli.go
```

## How to use
- type `cli -h` or `cli --help` to show detailed how-to-use
- Looking up for Name Server (eg. google.com)
```
cli ns --host google.com
```
- Resolving IP Address (eg. google.com)
```
cli ip --host google.com
```
- Looking up for CNAME / Canonical Name (eg. google.com)
```
cli cname --host google.com
```
- Looking up fo MX Records (eg. google.com)
```
cli mx --host google.com
```

## Issues
- Submit any issues to this [link](https://github.com/rikisan1993/go-website-lookup-tool/issues)

## Disclaimer
All the credits goes to [TutorialEdge](https://www.youtube.com/channel/UCwFl9Y49sWChrddQTD9QhRA) for their awesome tutorials
