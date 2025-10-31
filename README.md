# gomatrix

> A colorful terminal matrix rain effect with Hangul characters

A mesmerizing terminal screensaver inspired by `cmatrix`, but with a vibrant rainbow color scheme and beautiful Korean Hangul syllables cascading down your screen. Built with Go, Bubble Tea, and Lip Gloss.

---

## ✨ Features

- **Rainbow Colors**: Dynamic HSL-based color cycling creating a stunning rainbow effect
- **Hangul Characters**: Displays complete Korean syllables (가-힣) in elegant cascades
- **Smooth Animation**: High-performance rendering with optimized update cycles
- **Interactive**: Press `q` or `Ctrl+C` to exit
- **Terminal Native**: Full support for modern terminal emulators

---

## 📦 Installation

### Prerequisites

- Go 1.21 or later
- A terminal emulator with true color support (recommended)

### Build from Source

```bash
git clone https://github.com/guicybercode/gomatrix.git
cd gomatrix
go mod download
go build -o gomatrix
./gomatrix
```

### Install Globally

```bash
go install github.com/guicybercode/gomatrix@latest
```

---

## 🚀 Usage

Simply run the executable:

```bash
gomatrix
```

Press `q` or `Ctrl+C` to exit the program.

---

## 🎨 Technical Details

**Built with:**
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The fun, functional, stateful TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for nice terminal layouts

**Character Set:**
- Complete Hangul syllables only (U+AC00 to U+D7A3)
- No standalone jamo characters

**Color System:**
- HSL color space for smooth rainbow transitions
- Brightness fades along character trails
- Dynamic hue rotation based on position and time

---

## 📸 Screenshots

*Run the program to see the beautiful rainbow cascade of Korean characters!*

---

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">

**마태복음 28:20**

내가 너희에게 분부한 모든 것을 가르쳐 지키게 하라 볼지어다 내가 세상 끝날까지 너희와 항상 함께 있으리라 하시니라

</div>

