# prettyfmt
> Go fmt with colored prefix and timestamp

```shell
go get github.com/vikpe/prettyfmt
```

## Synopsis

Colors from [github.com/fatih/color](github.com/fatih/color)

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

