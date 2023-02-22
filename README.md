# go1cover

[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=asosnoviy_go1cover&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=asosnoviy_go1cover)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=asosnoviy_go1cover&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=asosnoviy_go1cover)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=asosnoviy_go1cover&metric=bugs)](https://sonarcloud.io/summary/new_code?id=asosnoviy_go1cover)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=asosnoviy_go1cover&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=asosnoviy_go1cover)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=asosnoviy_go1cover&metric=coverage)](https://sonarcloud.io/summary/new_code?id=asosnoviy_go1cover)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=asosnoviy_go1cover&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=asosnoviy_go1cover)

# Covered line collector for 1C:Enterprise

### Usage:
```bash
 go1cover start -s "./src" -u "localhost:1550"
 ~~run your test framework
 go1cover stop
 ```

 ### go1cover help - for more info

```
go1cover: line coverage collector for 1C:Enterprice through http debug server

Usage:
  go1cover [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  start       start coverage server
  stop        stop coverage server

Flags:
  -c, --config string   config file name (default ".go1cover.yaml")
  -h, --help            help for go1cover
  -p, --port string     coverage server port (default "4040")
  -v, --version         version for go1cover

Use "go1cover [command] --help" for more information about a command.
```
```
Usage:
  go1cover start [flags]

Flags:
  -b, --baseDir string    Project path (base dir for reports)
  -u, --debugger string   Debugger url (default "http://127.0.0.1:1550")
  -f, --format string     report format: lcov, generic (default "lcov, generic")
  -h, --help              help for start
  -s, --source string     Source directory to read from (default "./src")

Global Flags:
  -c, --config string   config file name (default ".go1cover.yaml")
  -p, --port string     coverage server port (default "4040")
```