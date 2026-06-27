# OdLang

A tree-walk interpreted programming language with Odia-inspired keywords, written in Go.

## Build

```bash
go build -o od ./cmd/od
```

## Run a program

```bash
./od run example/lexer_test.od
```

## REPL

```bash
./od repl
```

## Syntax quick reference

| Concept | Syntax |
|---|---|
| Print | `lekha("hello") \|` |
| Input | `dia(a >> b) \|` |
| Int variable | `sankhya x = 10 \|` |
| Float | `dasmik pi = 3.14 \|` |
| String | `sabda name = "Sameet" \|` |
| Bool | `satya ok = han \|` |
| Array | `krama arr(5, 0) \|` |
| If / else if / else | `jadi x > 0 { } nahele jadi x == 0 { } nahele { }` |
| For loop | `ghura sankhya i = 0 -> 5 \| i++ { }` |
| While loop | `jetebeleJain x > 0 { }` |
| Break / continue | `baharipade \|` / `chadide \|` |
| Function | `karya add(a sankhya, b sankhya) (sankhya) { deide (a + b) \| }` |
| Return | `deide (value) \|` |
| Try/catch | `chesta { } dhare { }` |
| Type cast | `sabda(x)` / `sankhya(x)` / `dasmik(x)` |
| Entry point | `karya aarambha() { }` |
| Statement end | `\|` (pipe character) |
| True / False / Null | `han` / `na` / `khali` |
| And / Or | `sahita` / `aau` |

## Example

```
karya aarambha() {
    sankhya x = 10 |
    jadi x > 5 {
        lekha("bada") |
    } nahele {
        lekha("chota") |
    }
}
```
