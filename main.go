package main

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	minSpeed = 2
	maxSpeed = 8
	trailLen = 25
)

type column struct {
	x         int
	chars     []rune
	positions []int
	speeds    []int
	counters  []int
}

type model struct {
	width   int
	height  int
	columns []column
	time    int
}

func initialModel() model {
	return model{
		columns: make([]column, 0),
		time:    0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.columns = make([]column, m.width)
		for i := range m.columns {
			m.columns[i] = newColumn(i, m.height)
		}
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
		return m, nil

	case tickMsg:
		m.time++
		for i := range m.columns {
			col := &m.columns[i]
			for j := range col.counters {
				col.counters[j]++
				if col.counters[j] >= col.speeds[j] {
					col.counters[j] = 0
					col.positions[j]++
					if col.positions[j] >= m.height+trailLen {
						col.positions[j] = -trailLen + rand.Intn(trailLen)
						col.chars[j] = randomHangul()
						col.speeds[j] = minSpeed + rand.Intn(maxSpeed-minSpeed+1)
					}
				}
			}
		}
		return m, tick()

	default:
		if len(m.columns) > 0 {
			return m, tick()
		}
		return m, nil
	}
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	grid := make([][]rune, m.height)
	for i := range grid {
		grid[i] = make([]rune, m.width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	for i := range m.columns {
		col := &m.columns[i]
		for j := range col.chars {
			pos := col.positions[j]
			if pos >= 0 && pos < m.height && col.x < m.width {
				grid[pos][col.x] = col.chars[j]
			}
			for k := 1; k <= trailLen; k++ {
				trailPos := pos - k
				if trailPos >= 0 && trailPos < m.height && col.x < m.width {
					if grid[trailPos][col.x] == ' ' {
						grid[trailPos][col.x] = randomHangul()
					}
				}
			}
		}
	}

	var lines []string
	for i := 0; i < m.height; i++ {
		line := ""
		for j := 0; j < m.width; j++ {
			char := grid[i][j]
			if char == ' ' {
				line += " "
				continue
			}

			distFromHead := trailLen
			for _, col := range m.columns {
				for _, pos := range col.positions {
					if pos == i && col.x == j {
						distFromHead = 0
						break
					}
					for k := 1; k <= trailLen; k++ {
						if pos-k == i && col.x == j {
							if k < distFromHead {
								distFromHead = k
							}
							break
						}
					}
				}
				if distFromHead == 0 {
					break
				}
			}

			hue := float64((m.time*2 + i*5 + j*3) % 360)
			brightness := 1.0 - (float64(distFromHead) / float64(trailLen) * 0.7)
			if brightness < 0.3 {
				brightness = 0.3
			}

			color := hslToColor(hue, 1.0, brightness*0.5)
			style := lipgloss.NewStyle().Foreground(color)
			line += style.Render(string(char))
		}
		lines = append(lines, line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func newColumn(x int, height int) column {
	cols := 5
	col := column{
		x:         x,
		chars:     make([]rune, cols),
		positions: make([]int, cols),
		speeds:    make([]int, cols),
		counters:  make([]int, cols),
	}

	for i := 0; i < cols; i++ {
		col.chars[i] = randomHangul()
		if height > 0 {
			col.positions[i] = -trailLen - rand.Intn(height/2)
		} else {
			col.positions[i] = -trailLen - rand.Intn(20)
		}
		col.speeds[i] = minSpeed + rand.Intn(maxSpeed-minSpeed+1)
		col.counters[i] = rand.Intn(col.speeds[i])
	}

	return col
}

func randomHangul() rune {
	return rune(0xAC00 + rand.Intn(0xD7A3-0xAC00+1))
}

func hslToColor(h, s, l float64) lipgloss.Color {
	if s == 0 {
		gray := uint8(l * 255)
		return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", gray, gray, gray))
	}

	var r, g, b float64
	var q, p float64

	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p = 2*l - q

	hNorm := h / 360.0
	r = hueToRGB(p, q, hNorm+1.0/3.0)
	g = hueToRGB(p, q, hNorm)
	b = hueToRGB(p, q, hNorm-1.0/3.0)

	return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x",
		uint8(r*255),
		uint8(g*255),
		uint8(b*255)))
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
