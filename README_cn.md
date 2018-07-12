# Name case

[![Build Status](https://travis-ci.org/wzshiming/namecase.svg?branch=master)](https://travis-ci.org/wzshiming/namecase)
[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/namecase)](https://goreportcard.com/report/github.com/wzshiming/namecase)
[![GoDoc](https://godoc.org/github.com/wzshiming/namecase?status.svg)](https://godoc.org/github.com/wzshiming/namecase)
[![GitHub license](https://img.shields.io/github/license/wzshiming/namecase.svg)](https://github.com/wzshiming/namecase/blob/master/LICENSE)
[![cover.run](https://cover.run/go/github.com/wzshiming/namecase.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fwzshiming%2Fnamecase)

各种命名风格转换

- [English](https://github.com/wzshiming/namecase/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/namecase/blob/master/README_cn.md)

## 安装

``` shell
go get -u -v github.com/wzshiming/namecase
```

## 用法

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

[API 文档](http://godoc.org/github.com/wzshiming/namecase)

## 许可证

软包根据MIT License。有关完整的许可证文本，请参阅[LICENSE](https://github.com/wzshiming/namecase/blob/master/LICENSE)。
