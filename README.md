# k64

> CLI that help you encode normal string in data fields of Kubernetes secret yaml to base64.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thetkpark/k64)![https://goreportcard.com/badge/github.com/thetkpark/k64](https://goreportcard.com/badge/github.com/thetkpark/k64)![https://img.shields.io/github/license/thetkpark/cscms](https://img.shields.io/github/license/thetkpark/cscms) 

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/you-didnt-ask-for-this.svg)](https://forthebadge.com)

When writing the Kubernetes secret declaration file in YAML format, the value of the data fields must be encoded using base64. I was annoy to just convert each secret one-by-one. Therefore, I built this CLI to help convert the normal secret value to base64 string.

![ezgif-6-45065a16c51c](https://user-images.githubusercontent.com/12962097/133083758-328f2b44-2fcb-48d9-8e83-e424c0647a0b.gif)

## Prerequisites

Install Go from https://golang.org

## Installation

```sh
$ go install github.com/thetkpark/k64
```

## Usage

To encode from string to base64

```sh
$ k64 encode <filename>
```

By default the output of that YAML file will be printed to stdout. You can use `-o` to write the output to another file or `-s` to save output to the save file.

```sh
$ k64 encode <filename> -s
```

```sh
$ k64 encode <filename> -o <filename>
```

To decode from base64 to string. The flag `-o` and `-s` can also be used here.

```sh
$ k64 decode <filename>
```
