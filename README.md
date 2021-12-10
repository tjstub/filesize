# filesize

## Description

`filesize` is just a small utility I wrote to display the file-sizes in units I cared about, rather than the bytes of `stat`. I do not happen to be good at math in my head, so this helps me.

## Installation

This is a pretty simple go program that does not use any external libraries. Just build and install the tool like you would any other go tool.

```sh
go build install
```

This was just to solve my simple problem, but suggestions and especially constructive feedback are welcome.

## Usage

```sh
filesize [-h] [-u <KB|MB|GB>] <file-names>
```

See `filesize -h` for more details.

## License

Licensed under GPLv3