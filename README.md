# Advent of Code preparator

This simple cli app set up a directory structure to solve a AoC day in Go.

It doesn't work on Windows since it doesn't have a keychain but works on MacOS and should work on Linux.

This app prevents multiple download of inputs from AoC by checking if the file already exists. I cannot guarantee that you
won't get banned tho so use at your own risk. (aka do you really trust me?)

Built with Go 1.17

## Usage

To set up for day 5 of 2021 in the current directory.

```bash
aocprep -y 2021 -d 5 -w .
```

## Build

```bash
go build
```

## Install from source

```bash
go install
```
