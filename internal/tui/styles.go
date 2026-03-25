package tui

import "github.com/charmbracelet/lipgloss"

// Color palette - Cyberpunk/Neon theme
var (
	// Primary colors
	neonPink   = lipgloss.Color("#FF006E")
	neonCyan   = lipgloss.Color("#00F5FF")
	neonGreen  = lipgloss.Color("#39FF14")
	neonYellow = lipgloss.Color("#FFE66D")
	neonOrange = lipgloss.Color("#FF9F1C")
	neonPurple = lipgloss.Color("#BF5AF2")

	// Neutral colors
	mediumBg   = lipgloss.Color("#1A1A2E")
	dimText    = lipgloss.Color("#4A4A6A")
	normalText = lipgloss.Color("#E4E4E7")
)

// Styles
var (
	// Title styles
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(neonCyan).
			Background(mediumBg).
			Padding(1, 4).
			MarginBottom(1)

	logoStyle = lipgloss.NewStyle().
			Foreground(neonPink).
			Bold(true)

	// Menu styles (set Width when rendering for responsive layout)
	menuStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(neonPurple).
			Padding(1, 2)

	menuItemStyle = lipgloss.NewStyle().
			Foreground(normalText).
			PaddingLeft(2)

	selectedMenuItemStyle = lipgloss.NewStyle().
				Foreground(neonCyan).
				Bold(true).
				PaddingLeft(2)

	// Algorithm info styles
	algoNameStyle = lipgloss.NewStyle().
			Foreground(neonPink).
			Bold(true).
			MarginBottom(1)

	algoDescStyle = lipgloss.NewStyle().
			Foreground(dimText).
			Italic(true)

	complexityStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true)

	barFillBG    = lipgloss.Color("#071b24")
	barCompareBG = lipgloss.Color("#241504")
	barSwapBG    = lipgloss.Color("#240818")
	barSortedBG  = lipgloss.Color("#071808")
	barPivotBG   = lipgloss.Color("#242006")

	// Visualization styles (wide bar glyphs + dim background in renderSortingBars)
	barStyle = lipgloss.NewStyle().
			Foreground(neonCyan).
			Background(barFillBG)

	comparingBarStyle = lipgloss.NewStyle().
				Foreground(neonOrange).
				Background(barCompareBG).
				Bold(true)

	swappingBarStyle = lipgloss.NewStyle().
				Foreground(neonPink).
				Background(barSwapBG).
				Bold(true)

	sortedBarStyle = lipgloss.NewStyle().
			Foreground(neonGreen).
			Background(barSortedBG)

	pivotBarStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Background(barPivotBG).
			Bold(true)

	barTrackStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#2f2f48"))
	// Slightly lighter top row of each bar (▓▓) for a simple highlight
	barTopCapStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#4fd4e8")).Background(barFillBG)
	barTopCapCmp    = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb347")).Background(barCompareBG).Bold(true)
	barTopCapSwap   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff6fb1")).Background(barSwapBG).Bold(true)
	barTopCapSorted = lipgloss.NewStyle().Foreground(lipgloss.Color("#7cff6a")).Background(barSortedBG)
	barTopCapPivot  = lipgloss.NewStyle().Foreground(lipgloss.Color("#fff1a8")).Background(barPivotBG).Bold(true)

	// Grid cell styles
	wallStyle = lipgloss.NewStyle().
			Foreground(dimText).
			Background(lipgloss.Color("#2D2D2D"))

	emptyStyle = lipgloss.NewStyle().
			Foreground(mediumBg)

	startStyle = lipgloss.NewStyle().
			Foreground(neonGreen).
			Bold(true)

	endStyle = lipgloss.NewStyle().
			Foreground(neonPink).
			Bold(true)

	visitedStyle = lipgloss.NewStyle().
			Foreground(neonPurple)

	frontierStyle = lipgloss.NewStyle().
			Foreground(neonOrange)

	currentStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true)

	pathStyle = lipgloss.NewStyle().
			Foreground(neonCyan).
			Bold(true)

	// Panel styles
	compactPanelStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(neonPurple).
				Padding(0, 1)

	// Control styles
	controlStyle = lipgloss.NewStyle().
			Foreground(dimText).
			MarginTop(1)

	keyStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true)

	// Status styles
	statusStyle = lipgloss.NewStyle().
			Foreground(neonGreen).
			Bold(true)

	descriptionStyle = lipgloss.NewStyle().
				Foreground(normalText).
				Italic(true).
				MarginTop(1)

	// Speed indicator
	speedStyle = lipgloss.NewStyle().
			Foreground(neonOrange).
			Bold(true)

	// Footer
	footerStyle = lipgloss.NewStyle().
			Foreground(dimText).
			MarginTop(2)
)

// Logo ASCII art
const logo = `
   █████╗ ██╗      ██████╗  ██████╗      █████╗ ██████╗ ███████╗███╗   ██╗ █████╗ 
  ██╔══██╗██║     ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██╔════╝████╗  ██║██╔══██╗
  ███████║██║     ██║  ███╗██║   ██║    ███████║██████╔╝█████╗  ██╔██╗ ██║███████║
  ██╔══██║██║     ██║   ██║██║   ██║    ██╔══██║██╔══██╗██╔══╝  ██║╚██╗██║██╔══██║
  ██║  ██║███████╗╚██████╔╝╚██████╔╝    ██║  ██║██║  ██║███████╗██║ ╚████║██║  ██║
  ╚═╝  ╚═╝╚══════╝ ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝
`

const smallLogo = `╔═══════════════════════════════════════╗
║     ALGO ARENA - Algorithm Visualizer ║
╚═══════════════════════════════════════╝`
