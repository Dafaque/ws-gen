WebSocket multi-language code generation tool

# Supported languages

| lang | client             | server                   | lib                      |
|------|--------------------|--------------------------|--------------------------|
| go   | :heavy_check_mark: | :heavy_check_mark:       | gorilla/websocket@v1.5.0 |
| dart | :heavy_check_mark: | :heavy_multiplication_x: | web_socket_channel@2.4.0 |

# Install
### With Go
`go install github.com/Dafaque/ws-gen`
### With Homebrew
`brew tap Dafaque/wsgen && brew install ws-gen`

# Usage
```
-g string
    which source files generate: client, server, all (default "all")
-l string
    target language (default "undefined")
-s string
    path to spec file (default "./wsgen.yml")
-v    show version
```
All language-specific params must be located in `config.wsgen.yml` file, in same directory with `wsgen.yml` file.

Required parameters are:
```
module: wsgen/examples # relative path for codegen
package: gen # codegen directory
```

Optional parameters:
```
root: ../ # path that will be prepended to "module" param
```
