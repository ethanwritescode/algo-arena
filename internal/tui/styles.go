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
	darkBg     = lipgloss.Color("#0D0D0D")
	mediumBg   = lipgloss.Color("#1A1A2E")
	lightBg    = lipgloss.Color("#16213E")
	dimText    = lipgloss.Color("#4A4A6A")
	normalText = lipgloss.Color("#E4E4E7")
	brightText = lipgloss.Color("#FFFFFF")
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

	// Menu styles
	menuStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(neonPurple).
			Padding(1, 2).
			Width(50)

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

	// Visualization styles
	barStyle = lipgloss.NewStyle().
			Foreground(neonCyan)

	comparingBarStyle = lipgloss.NewStyle().
				Foreground(neonOrange).
				Bold(true)

	swappingBarStyle = lipgloss.NewStyle().
				Foreground(neonPink).
				Bold(true)

	sortedBarStyle = lipgloss.NewStyle().
			Foreground(neonGreen)

	pivotBarStyle = lipgloss.NewStyle().
			Foreground(neonYellow).
			Bold(true)

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
	panelStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(neonPurple).
			Padding(1, 2)

	compactPanelStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(neonPurple).
				Padding(0, 1)

	infoPanelStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(neonCyan).
			Padding(0, 2).
			MarginTop(1)

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
║  ⚡ ALGO ARENA - Algorithm Visualizer ║
╚═══════════════════════════════════════╝`
