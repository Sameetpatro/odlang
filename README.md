# OdLang

A programming language with Odia-inspired keywords, built in Go.

![CI](https://github.com/Sameetpatro/odlang/actions/workflows/ci.yml/badge.svg)

---

## Install

### Windows (Scoop) — recommended
```powershell
scoop bucket add odlang https://github.com/Sameetpatro/scoop-odlang
scoop install odlang
```

### macOS (Homebrew)
```bash
brew tap Sameetpatro/odlang
brew install odlang
```

### Linux / Manual (all platforms)
Download the binary for your platform from [GitHub Releases](https://github.com/Sameetpatro/odlang/releases/latest).

**Linux:**
```bash
tar -xzf od_*_linux_amd64.tar.gz
sudo mv od /usr/local/bin/
```

**Windows (manual):**
Download `od_*_windows_amd64.zip`, extract, and add the folder to your PATH.

---

## Quick start

Create `hello.od`:
```
karya aarambha() {
    lekha("jai jagannath") ;
}
```

Run it:
```bash
od run hello.od
```

---

## CLI

```
od run <file.od>   Run an OdLang program
od repl            Start the interactive REPL
od version         Show version
od help            Show help
```

---

## Syntax

| Concept | Syntax |
|---|---|
| Print | `lekha("hello") ;` |
| Input | `dia(a >> b) ;` |
| Int | `sankhya x = 10 ;` |
| Float | `dasmik pi = 3.14 ;` |
| String | `sabda name = "Sameet" ;` |
| Bool | `satya ok = han ;` |
| Array | `krama arr(5, 0) ;` |
| If / else if / else | `jadi x > 0 { } nahele jadi x == 0 { } nahele { }` |
| For loop | `ghura sankhya i = 0 -> 5 ; i++ { }` |
| While loop | `jetebeleJain x > 0 { }` |
| Break / Continue | `baharipade ;` / `chadide ;` |
| Function | `karya add(a sankhya, b sankhya) (sankhya) { deide (a+b) ; }` |
| Return | `deide (value) ;` |
| Try / Catch | `chesta { } dhare { }` |
| Type cast | `sabda(x)` / `sankhya(x)` / `dasmik(x)` |
| Entry point | `karya aarambha() { }` |
| True / False / Null | `han` / `na` / `khali` |
| And / Or | `sahita` / `aau` |

## Keywords

| OdLang | Meaning |
|---|---|
| `lekha` | write / print |
| `dia` | give / input |
| `karya` | function |
| `aarambha` | beginning / main |
| `deide` | return |
| `jadi` | if |
| `nahele` | else |
| `ghura` | loop (for) |
| `jetebeleJain` | while |
| `baharipade` | break |
| `chadide` | continue |
| `sankhya` | number (int) |
| `dasmik` | decimal (float) |
| `sabda` | word (string) |
| `akshara` | character |
| `satya` | truth (bool) |
| `krama` | sequence (array) |
| `han` | yes (true) |
| `na` | no (false) |
| `khali` | empty (null) |
| `chesta` | try |
| `dhare` | catch |
| `anaa` | import |

---

## License

MIT © Sameet Patro
