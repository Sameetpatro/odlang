# OdLang

A tree-walk interpreted programming language with Odia-inspired keywords, written in Go.

## Build

```bash
go build -o od ./cmd/od
```

## Run a program

```bash
./od run example/hello.od
./od run example/class_test.od
```

Some programs read stdin — pipe input when needed:

```bash
echo "40 60" | ./od run example/lexer_test.od
```

## REPL

```bash
./od repl
```

## Examples

| File | What it covers |
|---|---|
| `example/hello.od` | Minimal entry point and print |
| `example/lexer_test.od` | Full language smoke test (types, loops, if/else, arrays, try/catch) |
| `example/class_test.od` | Classes (`sreni`), fields, and methods |

## Syntax quick reference

| Concept | Syntax |
|---|---|
| Print | `lekha("hello") ;` |
| Input | `dia(a >> b) ;` |
| Int variable | `sankhya x = 10 ;` |
| Float | `dasmik pi = 3.14 ;` |
| String | `sabda name = "Sameet" ;` |
| Bool | `satya ok = han ;` |
| Array | `krama arr(5, 0) ;` |
| Constant | `const PI = 3.14 ;` |
| Dynamic var | `nua x = khali ;` |
| Multi-assign | `s, p, q = misana(p, q) ;` |
| If / else if / else | `jadi x > 0 { } nahele jadi x == 0 { } nahele { }` |
| For loop | `ghura sankhya i = 0 -> 5 ; i++ { }` |
| While loop | `jetebeleJain x > 0 { }` |
| Break / continue | `baharipade ;` / `chadide ;` |
| Function | `karya add(a sankhya, b sankhya) (sankhya) { deide (a + b) ; }` |
| Function (no return type) | `karya aarambha() { }` |
| Return | `deide (value) ;` |
| Class | `sreni Point { sankhya x ; karya sum() (sankhya) { deide (x) ; } }` |
| New instance | `nua p = Point() ;` |
| Method call | `p.init(3, 4) ;` |
| Field access | `p.x` |
| Import | `anaa fmt ;` (parsed; module loading not yet implemented) |
| Try/catch | `chesta { } dhare { }` |
| Type cast | `sabda(x)` / `sankhya(x)` / `dasmik(x)` |
| Entry point | `karya aarambha() { }` |
| Statement end | `;` (semicolon) |
| True / False / Null | `han` / `na` / `khali` |
| And / Or | `sahita` / `aau` |

## Example

```od
sreni Point {
    sankhya x ;
    sankhya y ;

    karya init(ax sankhya, ay sankhya) () {
        x = ax ;
        y = ay ;
    }

    karya sum() (sankhya) {
        deide (x + y) ;
    }
}

karya aarambha() {
    nua p = Point() ;
    p.init(3, 4) ;
    lekha("sum = " + sabda(p.sum())) ;
}
```
