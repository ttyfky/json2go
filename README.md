# json2go

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

json2go is a command line tool to generate Go struct from JSON file.

# Usecase

This tool is useful when creating Go struct models for Web APIs by collecting JSON payload as sample data sets.

Making models for numbers of API payloads is time consuming. This tool assist the preparation and developers can focus
on creation of core logic.

## Models require follow up

Generated code is not intended to be used as it is.

JSON cannot describe detail such as precise data type or underlying data models of each field. Therefore generated code
is not expected to be used and just being referenced, bu additional modification is expected as following points.

* Auto mappable data type between JSON and Go is limited.
* If values in JSON is not set the value type is set as `interface{}`
* Generated name pattern is all CamelCase. ID, HTML, or other abbreviation is not considered.

# Getting started

## Download

### Binary

Runnable binary file is put in the release page. The tool is usable by downloading the package.



### Go src

Download source by command below and build.

```console
go get https://github.com/ttyfky/json2go
```

## Use

`json2go` command needs argument of input path.

Tha path can be:

* Path to a JSON file
    * A Model is generated for single JSON file.

* Path to a directory
    * Models are generated for all JSON files in the provided directory.

```shell
json2go [options] $PATH_TO_DIR_OR_FILE
## Options
```

* `-output`: output is a path to put generated files. If this is not specified the generated models are printed in
  stdout.

* `-pkg`: package name of generated models.

## Example

Sample input JSON `person.json`

```json
{
  "id": 1,
  "first_name": "Taro",
  "last_name": "Tanaka",
  "address": {
    "zip": "123-4567",
    "country": "Japan",
    "city": "Tokyo"
  },
  "brothers": [
    {
      "first_name": "Jiro",
      "last_name": "Tanaka",
      "relation": "younger"
    },
    {
      "first_name": "Hanako",
      "last_name": "Tanaka",
      "relation": "younger"
    }
  ],
  "age": 30,
  "is_married": false
}
```

Command

```shell
 json2go -pkg model input/member.json
```

Output

```go
package model

type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Zip     string `json:"zip,omitempty"`
}
type Brother struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Relation  string `json:"relation,omitempty"`
}
type Member struct {
	Address   *Address   `json:"address,omitempty"`
	Age       int64      `json:"age,omitempty"`
	Brothers  []*Brother `json:"brothers,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	Id        int64      `json:"id,omitempty"`
	IsMarried bool       `json:"is_married,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
}

```
