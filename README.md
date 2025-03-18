# Language Detection Benchmark: Whatlang vs Lingua

This project benchmarks the performance and accuracy of two language detection libraries:
- [Whatlang](https://github.com/abadojack/whatlanggo)
- [lingua](https://github.com/pemistahl/lingua)

## Benchmark Results

| Library | Average Time | Accuracy |
|---------|--------------|-----------|
| Whatlang | 0.26 ms | 90.79% |
| Lingua (All) | 23.60 ms | 94.74% |
| Lingua (65) | 0.45 ms | 94.74% |
| Lingua (10) | 0.12 ms | 52.63% |

## Setup

```bash
go get -u github.com/abadojack/whatlanggo
go get -u github.com/pemistahl/lingua-golang
```


## Usage

```bash
go run benchmark-whatlang-vs-lingua
```