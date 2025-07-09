# Go‑Reloaded 🎯

A simple **text completion / editing / auto‑correction** tool written in Go using only the standard library.

---

## ⚙️ Features

- **Hex → Decimal**: Converts hexadecimal numbers before `(hex)` to decimal.
- **Bin → Decimal**: Converts binary numbers before `(bin)` to decimal.
- **Uppercase / Lowercase / Capitalize**: Supports `(up)`, `(low)`, `(cap)` and also multi-word modifiers like `(up, 2)`.
- **Punctuation Formatting**: Properly spaces punctuation `. , ! ? : ;`, handles grouped punctuation like `…` or `!?`.
- **Apostrophes**: Formats `'word'` or `’multiple words’` correctly.
- **Grammar Fix**: Swaps `a` → `an` before vowels or an ‘h’.

---

## 🧪 Getting Started

### Prerequisites

- Go (version 1.20+ recommended)

### Installation

```bash
git clone https://github.com/AhmedTalbii/Go-Reloaded.git
cd Go-Reloaded
go build -o go-reloaded main.go
```

This builds the `go-reloaded` binary.

---

## 🚀 Usage

```bash
./go-reloaded sample.txt result.txt
```

- `sample.txt` – input file with text to modify
- `result.txt` – output file with applied transformations

### Example

Input (`sample.txt`):

```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom,...
```

Run:

```bash
./go-reloaded sample.txt result.txt
```

Output (`sample.txt` → `result.txt`):

```
It was the best of times, it was the worst of TIMES,...
```

---

## 📁 Project Structure

```
Go-Reloaded/
├── main.go         # Entry point
├── go.mod
└── helpers/        # Utility functions
    ├── aToAn.go
    ├── punctuation.go
    ├── toDecimal.go
    └── ...
```

---

## 🧩 Future Enhancements

- Add **unit tests** for all helper modules.
- Provide **config flags** (e.g., toggle specific transformations).
- Print help/version info with `--help` / `--version`.

---

## 🤝 Contributing

1. Fork the repo
2. Create a feature branch (`git checkout -b feature/XYZ`)
3. Commit your changes (`git commit -m "Add XYZ"`)
4. Push and open a Pull Request

Suggestions and improvements are always welcome!

---

## 📄 License

This project is licensed under the **MIT License**. See [LICENSE](./LICENSE) for details.
