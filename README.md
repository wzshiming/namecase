# Name case

[![Build Status](https://travis-ci.org/wzshiming/namecase.svg?branch=master)](https://travis-ci.org/wzshiming/namecase)
[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/namecase)](https://goreportcard.com/report/github.com/wzshiming/namecase)
[![GoDoc](https://godoc.org/github.com/wzshiming/namecase?status.svg)](https://godoc.org/github.com/wzshiming/namecase)
[![GitHub license](https://img.shields.io/github/license/wzshiming/namecase.svg)](https://github.com/wzshiming/namecase/blob/master/LICENSE)
[![cover.run](https://cover.run/go/github.com/wzshiming/namecase.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fwzshiming%2Fnamecase)

Various naming styles change.

- [English](https://github.com/wzshiming/namecase/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/namecase/blob/master/README_cn.md)

## Install

``` shell
go get -u -v github.com/wzshiming/namecase
```

## Usage

|               |                 |
| ------------- | --------------- |
| ToUpperSpace  | all to XX YY ZZ |
| ToLowerSpace  | all to xx yy zz |
| ToUpperStrike | all to XX-YY-ZZ |
| ToLowerStrike | all to xx-yy-zz |
| ToUpperSnake  | all to XX_YY_ZZ |
| ToLowerSnake  | all to xx_yy_zz |
| ToPascal      | all to XxYyZz   |
| ToUpperHump   | all to XxYyZz   |
| ToCamel       | all to xxYyZz   |
| ToLowerHump   | all to xxYyZz   |

[API Documentation](http://godoc.org/github.com/wzshiming/namecase)

## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/namecase/blob/master/LICENSE) for the full license text.
