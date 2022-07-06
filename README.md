# prettyfmt [![Test](https://github.com/vikpe/prettyfmt/actions/workflows/test.yml/badge.svg)](https://github.com/vikpe/prettyfmt/actions/workflows/test.yml) [![codecov](https://codecov.io/gh/vikpe/prettyfmt/branch/main/graph/badge.svg?token=sc5Hd7M4wv)](https://codecov.io/gh/vikpe/prettyfmt)
> Go fmt with colored prefix and timestamp

```shell
go get github.com/vikpe/prettyfmt
```

## Synopsis

Colors from [github.com/fatih/color](https://github.com/fatih/color)

```
New(
  prefix          string,
  prefixColor     color.Attribute, 
  timestampFormat string, 
  timestampColor  color.Attribute
)
```

## Usage
```go
pfmt := prettyfmt.New("chatbot", color.FgHiBlue, "15:04:05", color.FgWhite)
pfmt.Println("start")
pfmt.Printfln("connected as %s", "vikpebot")
```

![image](https://user-images.githubusercontent.com/1616817/177564266-297c02a7-d5be-40b8-aac0-d97b3829830c.png)

