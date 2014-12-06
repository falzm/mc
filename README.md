mc – basic CLI for memcache
===========================

`mc` is a small utility to manage keys stored in a memcache server. It is very basic, but handy to handle values from a shell.

Installation
------------

```
go get github.com/falzm/mc
```

Usage
-----

```
$ mc get bob
NOT FOUND
$ mc set bob kelso
OK
$ mc get bob
kelso
$ mc replace bob marley
OK
$ mc get bob
marley
$ mc delete bob
OK
$ mc get bob
NOT FOUND
$ mc set key1 1
OK
$ mc set key2 2
OK
$ mc set key3 3
OK
$ mc get key1
1
$ mc get key2
2
$ mc get key3
3
$ mc flush
OK
$ mc get key1
NOT FOUND
$ mc get key2
NOT FOUND
$ mc get key3
NOT FOUND
```

Global usage documentation:

```
$ mc --help
usage: mc [--version] [--help] <command> [<args>]

Available commands are:
    delete     delete a key
    flush      flush all keys
    get        get the value of a key
    replace    replace the value to a key if it exists
    set        set a value to a key
    touch      update the expiry of a key
```

Commands usage documentation:

```
$ mc --help get
Usage: mc get [options] <key>

  Get the value associated to <key> on memcached server

Options:

  -serverHost=HOST  memcached server host (default: 127.0.0.1)
  -serverPort=PORT  memcached server port (default: 11211)
```

```
$ mc --help set
Usage: mc set [options] <key> <value>

  Set <key> to <value> on memcached server

Options:

  -serverHost=HOST  memcached server host (default: 127.0.0.1)
  -serverPort=PORT  memcached server port (default: 11211)
  -ttl=N            key expiration time in seconds (default: 0, i.e. no expiration)
```

License
--------

The MIT License (MIT)

Copyright (c) 2014 Marc Falzon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

