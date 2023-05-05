:bangbang: **This software is in early alpha stage and not recomended to anyone.** :bangbang:

Check [this](https://github.com/users/Dafaque/projects/1) page and help me!

-----


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
`brew tap Dafaque/ws-gen && brew install ws-gen`

# Usage
```
-c string
    path to config file (default "wsgen.config.yml")
-g string
    which source files generate: client, server, all (default "all")
-l string
    target language (default "undefined")
-s string
    path to spec file (default "wsgen.spec.yml")
-v    show version
```

# Data primitives
```
* int8
* int16
* int32
* int64
* int
* uint8
* uint16
* uint32
* uint64
* uint
* float
* float32
* float64
* string
* bool
```
Each type have two modifiers combinable:
1. prefix `...`, which means array (list)
2. suffix `?`, which means optional (nullable) type

# Spec reference
```
# message packing format
encoding: json
# init section describes GET variables
# that will be passed before WebSocket
# connection start to handle messages.
init:
  params: 
    - name: chatId # well be converted in snake_case in URL
    - name: invisible
      optional: true # no error if empty
# enums section describes, well, enums.
# values will be packed as indexes
# 0 index reserved for "undefined"
# you can reference enums by typing $ and "name" of enum
enums:
  - name: event
    values:
      - entered
      - leaved
# aliases is not a spec section
# i use it here just for "caching" some fields
# that will be used more than once
# you may read more about YAML aliases 
# and anchors and be free to use them here.
aliases: 
  id: &id
    name: id
    type: int64
# messages section is where all boilerplate lives
# every message will contain $msg_idx field
# by which Handler will call On{$message.name} method
# it is better to use generated message constructors
# wich will assign $msg_idx for you 
messages:
  - name: textMessage # use camelCase, please
    fields:
      - *id # alias
      - name: content
        type: string? # optional (nullable) string
      - name: arrayFieldExample
        type: ...int64 # array (list) of integers
      - name: arrayOptionalFieldExample
        type: ...float64? # array (list) of otional (nullable) floats
  - name: chatEvent
    fields:
      - *id
      - name: event
        type: $event # enum reference
      - name: testSnakeCaseConvertor # camelCase, please
        type: float
```

# Config
All project-specific params must be located in `wsgen.config.yml` file.

Required parameters:
```
module: wsgen-project # name of you app or module. Used for imports
package: examples/gen # relative path for codegen
```

Optional parameters:
```
root: ../ # path that will be prepended to "package" param
```

# Contribute
1. Create an issue
1. Clone
1. Create branch
1. Hack-hack-hack
1. Open pull requests
1. Wait for approval
1. ???
