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

## Adding Test Cases

Test cases are stored in `test_samples.txt` in the following format:
```
language_code|text_sample
```

For example:
```
en|The quick brown fox jumps over the lazy dog.
es|El veloz zorro marrón salta sobre el perro perezoso.
```

The language code should be in ISO 639-1 format (e.g., `en` for English, `es` for Spanish, etc.). Each test case should be on a new line, with the language code and text sample separated by a pipe character (`|`).

Currently supported language codes:
- `en` - English
- `es` - Spanish
- `fr` - French
- `de` - German
- `it` - Italian
- `pt` - Portuguese
- `nl` - Dutch
- `ru` - Russian
- `ja` - Japanese
- `zh` - Chinese
- `ko` - Korean
- `ar` - Arabic
- `hi` - Hindi
- `tr` - Turkish
- `vi` - Vietnamese
- `sv` - Swedish
- `pl` - Polish
- `fi` - Finnish
- `gr` - Greek
