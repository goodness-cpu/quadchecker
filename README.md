# quad-checker

A command-line tool written in Go that checks whether an ASCII figure is a valid quadrilateral — and if so, tells you exactly what kind.

## What it does

You pass an ASCII shape through the command line, and `quad-checker` analyzes it and tells you what kind of quadrilateral it is — or tells you it isn't one at all.

Possible results include:

- **Square**
- **Rectangle**
- **Rhombus**
- **Parallelogram**
- **Trapezoid**
- **Invalid** — if the figure doesn't form a valid quadrilateral

## Usage

```bash
go run . "<your ASCII figure>"
```

### Examples

```bash
go run . "+---------+
|         |
|         |
+---------+"
# Output: Rectangle

go run . "+----+
|    |
|    |
+----+"
# Output: Square
```

## How it works

The program reads the ASCII figure passed as a command-line argument, parses its structure, and applies geometric logic to classify the shape based on its dimensions and proportions.

No external libraries are used — just Go's standard library.

## Getting started

**Prerequisites:** Go 1.18 or later



Or build a binary:

```bash
go build -o quad-checker
./quad-checker "<figure>"
```

## Built with

- [Go](https://golang.org/) — standard library only

## Author

**Goodness Omotine** — [@goodness-cpu](https://github.com/goodness-cpu)

---
