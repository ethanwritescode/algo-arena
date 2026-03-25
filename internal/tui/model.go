package tui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethanwritescode/algo-arena/internal/algorithms/pathfinding"
	"github.com/ethanwritescode/algo-arena/internal/algorithms/sorting"
)

// View states
type viewState int

const (
	viewMenu viewState = iota
	viewSortingMenu
	viewPathfindingMenu
	viewSortingVis
	viewPathfindingVis
	viewAbout
)

// Speed settings
type speedSetting int

const (
	speedSlow speedSetting = iota
	speedNormal
	speedFast
)

func (s speedSetting) String() string {
	switch s {
	case speedSlow:
		return "🐢 Slow"
	case speedNormal:
		return "🐇 Normal"
	case speedFast:
		return "⚡ Fast"
	}
	return "Normal"
}

func (s speedSetting) Duration() time.Duration {
	switch s {
	case speedSlow:
		return 500 * time.Millisecond
	case speedNormal:
		return 50 * time.Millisecond
	case speedFast:
		return 5 * time.Millisecond
	}
	return 50 * time.Millisecond
}

// Model is the main application model
type Model struct {
	width      int
	height     int
	view       viewState
	menuCursor int
	speed      speedSetting
	paused     bool

	// Sorting visualization
	sortingAlgo  *sorting.Algorithm
	sortingStep  int
	sortingArray []int

	// Pathfinding visualization
	pathfindingAlgo *pathfinding.Algorithm
	pathfindingStep int
	pathfindingGrid *pathfinding.Grid
}

// Menu items
var mainMenuItems = []string{
	"🔢 Sorting Algorithms",
	"🗺️  Pathfinding Algorithms",
	"❓ About",
	"🚪 Quit",
}

var sortingMenuItems = []string{
	"🫧 Bubble Sort",
	"📌 Selection Sort",
	"📥 Insertion Sort",
	"🐚 Shell Sort",
	"⚡ Quick Sort",
	"🔀 Merge Sort",
	"🏔️  Heap Sort",
	"← Back",
}

var pathfindingMenuItems = []string{
	"🌊 Breadth-First Search (BFS)",
	"🌲 Depth-First Search (DFS)",
	"📍 Dijkstra's Algorithm",
	"⭐ A* Search Algorithm",
	"← Back",
}

// NewModel creates a new model
func NewModel() Model {
	return Model{
		width:  120,
		height: 40,
		view:   viewMenu,
		speed:  speedNormal,
	}
}

// Tick message for animation
type tickMsg time.Time

func tickCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case tickMsg:
		return m.handleTick()
	}

	return m, nil
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "q":
		if m.view == viewMenu {
			return m, tea.Quit
		}
		return m.handleBack()

	case "up", "k":
		if m.menuCursor > 0 {
			m.menuCursor--
		}
		return m, nil

	case "down", "j":
		maxItems := m.getMaxMenuItems()
		if m.menuCursor < maxItems-1 {
			m.menuCursor++
		}
		return m, nil

	case "enter", " ":
		return m.handleSelect()

	case "r":
		return m.handleReset()

	case "p":
		m.paused = !m.paused
		if !m.paused && (m.view == viewSortingVis || m.view == viewPathfindingVis) {
			return m, tickCmd(m.speed.Duration())
		}
		return m, nil

	case "1":
		m.speed = speedSlow
		return m, nil

	case "2":
		m.speed = speedNormal
		return m, nil

	case "3":
		m.speed = speedFast
		return m, nil

	case "left", "h":
		return m.handleStepBack()

	case "right", "l":
		return m.handleStepForward()

	case "esc":
		return m.handleBack()
	}

	return m, nil
}

func (m Model) getMaxMenuItems() int {
	switch m.view {
	case viewMenu:
		return len(mainMenuItems)
	case viewSortingMenu:
		return len(sortingMenuItems)
	case viewPathfindingMenu:
		return len(pathfindingMenuItems)
	}
	return 0
}

func (m Model) handleSelect() (tea.Model, tea.Cmd) {
	switch m.view {
	case viewMenu:
		switch m.menuCursor {
		case 0: // Sorting
			m.view = viewSortingMenu
			m.menuCursor = 0
		case 1: // Pathfinding
			m.view = viewPathfindingMenu
			m.menuCursor = 0
		case 2: // About
			m.view = viewAbout
			return m, nil
		case 3: // Quit
			return m, tea.Quit
		}

	case viewSortingMenu:
		if m.menuCursor == len(sortingMenuItems)-1 {
			m.view = viewMenu
			m.menuCursor = 0
			return m, nil
		}
		return m.startSortingVisualization()

	case viewPathfindingMenu:
		if m.menuCursor == len(pathfindingMenuItems)-1 {
			m.view = viewMenu
			m.menuCursor = 1
			return m, nil
		}
		return m.startPathfindingVisualization()
	}

	return m, nil
}

func (m Model) startSortingVisualization() (tea.Model, tea.Cmd) {
	// Adjust array size based on terminal width (each element takes ~2 chars)
	arraySize := min(max((m.width-10)/3, 15), 50)
	m.sortingArray = sorting.GenerateRandomArray(arraySize)

	switch m.menuCursor {
	case 0:
		m.sortingAlgo = sorting.BubbleSort(m.sortingArray)
	case 1:
		m.sortingAlgo = sorting.SelectionSort(m.sortingArray)
	case 2:
		m.sortingAlgo = sorting.InsertionSort(m.sortingArray)
	case 3:
		m.sortingAlgo = sorting.ShellSort(m.sortingArray)
	case 4:
		m.sortingAlgo = sorting.QuickSort(m.sortingArray)
	case 5:
		m.sortingAlgo = sorting.MergeSort(m.sortingArray)
	case 6:
		m.sortingAlgo = sorting.HeapSort(m.sortingArray)
	}

	m.sortingStep = 0
	m.paused = false
	m.view = viewSortingVis
	return m, tickCmd(m.speed.Duration())
}

func (m Model) startPathfindingVisualization() (tea.Model, tea.Cmd) {
	// Adjust grid size based on terminal dimensions
	// Account for UI chrome (header, status, legend, controls, borders)
	// Use odd dimensions for proper maze generation
	gridWidth := min(max((m.width-10)/2, 15), 41)
	gridHeight := min(max(m.height-15, 9), 21)
	// Ensure odd dimensions for maze algorithm
	if gridWidth%2 == 0 {
		gridWidth--
	}
	if gridHeight%2 == 0 {
		gridHeight--
	}
	m.pathfindingGrid = pathfinding.NewGrid(gridWidth, gridHeight, 0.25)

	switch m.menuCursor {
	case 0:
		m.pathfindingAlgo = pathfinding.BFS(m.pathfindingGrid)
	case 1:
		m.pathfindingAlgo = pathfinding.DFS(m.pathfindingGrid)
	case 2:
		m.pathfindingAlgo = pathfinding.Dijkstra(m.pathfindingGrid)
	case 3:
		m.pathfindingAlgo = pathfinding.AStar(m.pathfindingGrid)
	}

	m.pathfindingStep = 0
	m.paused = false
	m.view = viewPathfindingVis
	return m, tickCmd(m.speed.Duration())
}

func (m Model) handleTick() (tea.Model, tea.Cmd) {
	if m.paused {
		return m, nil
	}

	switch m.view {
	case viewSortingVis:
		if m.sortingAlgo != nil && m.sortingStep < len(m.sortingAlgo.Steps)-1 {
			m.sortingStep++
			return m, tickCmd(m.speed.Duration())
		}

	case viewPathfindingVis:
		if m.pathfindingAlgo != nil && m.pathfindingStep < len(m.pathfindingAlgo.Steps)-1 {
			m.pathfindingStep++
			return m, tickCmd(m.speed.Duration())
		}
	}

	return m, nil
}

func (m Model) handleReset() (tea.Model, tea.Cmd) {
	switch m.view {
	case viewSortingVis:
		return m.startSortingVisualization()
	case viewPathfindingVis:
		return m.startPathfindingVisualization()
	}
	return m, nil
}

func (m Model) handleStepBack() (tea.Model, tea.Cmd) {
	if !m.paused {
		return m, nil
	}
	switch m.view {
	case viewSortingVis:
		if m.sortingStep > 0 {
			m.sortingStep--
		}
	case viewPathfindingVis:
		if m.pathfindingStep > 0 {
			m.pathfindingStep--
		}
	}
	return m, nil
}

func (m Model) handleStepForward() (tea.Model, tea.Cmd) {
	if !m.paused {
		return m, nil
	}
	switch m.view {
	case viewSortingVis:
		if m.sortingAlgo != nil && m.sortingStep < len(m.sortingAlgo.Steps)-1 {
			m.sortingStep++
		}
	case viewPathfindingVis:
		if m.pathfindingAlgo != nil && m.pathfindingStep < len(m.pathfindingAlgo.Steps)-1 {
			m.pathfindingStep++
		}
	}
	return m, nil
}

func (m Model) handleBack() (tea.Model, tea.Cmd) {
	switch m.view {
	case viewSortingVis:
		m.view = viewSortingMenu
	case viewPathfindingVis:
		m.view = viewPathfindingMenu
	case viewSortingMenu, viewPathfindingMenu, viewAbout:
		m.view = viewMenu
		m.menuCursor = 0
	}
	return m, nil
}

// View renders the UI
func (m Model) View() string {
	switch m.view {
	case viewMenu:
		return m.viewMainMenu()
	case viewSortingMenu:
		return m.viewSortingMenu()
	case viewPathfindingMenu:
		return m.viewPathfindingMenu()
	case viewSortingVis:
		return m.viewSortingVisualization()
	case viewPathfindingVis:
		return m.viewPathfindingVisualization()
	case viewAbout:
		return m.viewAbout()
	}
	return ""
}

func (m Model) viewMainMenu() string {
	var b strings.Builder

	// Use small logo for smaller terminals
	if m.height < 30 {
		b.WriteString(logoStyle.Render(smallLogo))
	} else {
		b.WriteString(logoStyle.Render(logo))
	}
	b.WriteString("\n")

	// Subtitle (skip for very small terminals)
	if m.height >= 20 {
		subtitle := lipgloss.NewStyle().
			Foreground(dimText).
			Italic(true).
			Render("  Watch algorithms come to life in your terminal")
		b.WriteString(subtitle)
		b.WriteString("\n")
	}
	b.WriteString("\n")

	// Menu
	menuContent := m.renderMenu(mainMenuItems)
	b.WriteString(menuStyle.Render(menuContent))
	b.WriteString("\n")

	// Controls
	controls := controlStyle.Render(
		keyStyle.Render("↑/↓") + " Navigate  " +
			keyStyle.Render("Enter") + " Select  " +
			keyStyle.Render("q") + " Quit",
	)
	b.WriteString(controls)

	// Footer (only for larger terminals)
	if m.height >= 25 {
		footer := footerStyle.Render("\n  Made with 💜 using Bubble Tea • github.com/ethanwritescode/algo-arena")
		b.WriteString(footer)
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, b.String())
}

func (m Model) viewSortingMenu() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("🔢 SORTING ALGORITHMS"))
	b.WriteString("\n")

	// Description (only for larger terminals)
	if m.height >= 25 {
		desc := lipgloss.NewStyle().
			Foreground(dimText).
			Width(50).
			Render("Visualize how different sorting algorithms organize data. Watch comparisons, swaps, and see the array transform in real-time.")
		b.WriteString(desc)
		b.WriteString("\n")
	}
	b.WriteString("\n")

	// Menu
	menuContent := m.renderMenu(sortingMenuItems)
	b.WriteString(menuStyle.Render(menuContent))
	b.WriteString("\n")

	// Controls
	controls := controlStyle.Render(
		keyStyle.Render("↑/↓") + " Navigate  " +
			keyStyle.Render("Enter") + " Select  " +
			keyStyle.Render("Esc") + " Back",
	)
	b.WriteString(controls)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, b.String())
}

func (m Model) viewPathfindingMenu() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("🗺️ PATHFINDING ALGORITHMS"))
	b.WriteString("\n")

	// Description (only for larger terminals)
	if m.height >= 25 {
		desc := lipgloss.NewStyle().
			Foreground(dimText).
			Width(50).
			Render("Watch pathfinding algorithms navigate through a maze. See how different strategies explore and find the optimal path.")
		b.WriteString(desc)
		b.WriteString("\n")
	}
	b.WriteString("\n")

	// Menu
	menuContent := m.renderMenu(pathfindingMenuItems)
	b.WriteString(menuStyle.Render(menuContent))
	b.WriteString("\n")

	// Controls
	controls := controlStyle.Render(
		keyStyle.Render("↑/↓") + " Navigate  " +
			keyStyle.Render("Enter") + " Select  " +
			keyStyle.Render("Esc") + " Back",
	)
	b.WriteString(controls)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, b.String())
}

func (m Model) viewAbout() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("❓ ABOUT ALGO ARENA"))
	b.WriteString("\n\n")

	content := lipgloss.NewStyle().Foreground(normalText).Width(56)

	b.WriteString(content.Render("Algo Arena is an interactive terminal visualizer for sorting and pathfinding algorithms. Watch algorithms come to life step by step."))
	b.WriteString("\n\n")

	b.WriteString(algoNameStyle.Render("Sorting Algorithms"))
	b.WriteString("\n")
	b.WriteString(content.Render("Bubble, Selection, Insertion, Shell, Quick, Merge, and Heap sort — with real-time comparisons and swap counts."))
	b.WriteString("\n\n")

	b.WriteString(algoNameStyle.Render("Pathfinding Algorithms"))
	b.WriteString("\n")
	b.WriteString(content.Render("BFS, DFS, Dijkstra, and A* — navigating through procedurally generated mazes with live frontier and visited cell tracking."))
	b.WriteString("\n\n")

	controls := controlStyle.Render(
		keyStyle.Render("Esc") + " Back",
	)
	b.WriteString(controls)

	if m.height >= 25 {
		footer := footerStyle.Render("\n  Built with Bubble Tea + Lipgloss • github.com/ethanwritescode/algo-arena")
		b.WriteString(footer)
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, b.String())
}

func (m Model) renderMenu(items []string) string {
	var b strings.Builder

	for i, item := range items {
		cursor := "  "
		style := menuItemStyle
		if i == m.menuCursor {
			cursor = "▸ "
			style = selectedMenuItemStyle
		}
		b.WriteString(style.Render(cursor + item))
		b.WriteString("\n")
	}

	return b.String()
}

func (m Model) viewSortingVisualization() string {
	if m.sortingAlgo == nil || len(m.sortingAlgo.Steps) == 0 {
		return "Loading..."
	}

	step := m.sortingAlgo.Steps[m.sortingStep]
	var b strings.Builder

	// Header (compact for small terminals)
	if m.height >= 25 {
		header := lipgloss.JoinHorizontal(lipgloss.Top,
			algoNameStyle.Render(m.sortingAlgo.Name),
			"  ",
			complexityStyle.Render("Time: "+m.sortingAlgo.TimeComplex),
			"  ",
			complexityStyle.Render("Space: "+m.sortingAlgo.SpaceComplex),
		)
		b.WriteString(header)
		b.WriteString("\n")
		b.WriteString(algoDescStyle.Render(m.sortingAlgo.Description))
		b.WriteString("\n")
	} else {
		b.WriteString(algoNameStyle.Render(m.sortingAlgo.Name))
		b.WriteString("\n")
	}

	// Visualization
	vis := m.renderSortingBars(step)
	b.WriteString(compactPanelStyle.Render(vis))
	b.WriteString("\n")

	// Status with stats
	moveLabel := m.sortingAlgo.MoveStatLabel
	if moveLabel == "" {
		moveLabel = "Swp"
	}
	progress := fmt.Sprintf("Step %d/%d", m.sortingStep+1, len(m.sortingAlgo.Steps))
	statsLine := lipgloss.JoinHorizontal(lipgloss.Top,
		statusStyle.Render(progress),
		"  ",
		speedStyle.Render(m.speed.String()),
		"  ",
		lipgloss.NewStyle().Foreground(neonCyan).Render(fmt.Sprintf("Cmp: %d", step.Comparisons)),
		"  ",
		lipgloss.NewStyle().Foreground(neonPink).Render(fmt.Sprintf("%s: %d", moveLabel, step.Swaps)),
		"  ",
		m.getPauseStatus(),
	)
	b.WriteString(statsLine)

	// Description (skip for small terminals)
	if m.height >= 22 {
		b.WriteString("\n")
		b.WriteString(descriptionStyle.Render("▸ " + step.Description))
	}
	b.WriteString("\n")

	// Legend (only for larger terminals)
	if m.height >= 25 {
		legend := m.renderSortingLegend()
		b.WriteString(legend)
		b.WriteString("\n")
	}

	controls := controlStyle.Render(
		keyStyle.Render("1/2/3") + " Speed  " +
			keyStyle.Render("p") + " Pause  " +
			keyStyle.Render("←/→") + " Step  " +
			keyStyle.Render("r") + " Reset  " +
			keyStyle.Render("Esc") + " Back",
	)
	b.WriteString(controls)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, b.String())
}

func (m Model) renderSortingBars(step sorting.Step) string {
	if len(step.Array) == 0 {
		return ""
	}

	maxVal := 0
	for _, v := range step.Array {
		if v > maxVal {
			maxVal = v
		}
	}

	// Pre-compute state sets for O(1) lookups instead of linear scans
	swapSet := make(map[int]bool, len(step.Swapping))
	for _, idx := range step.Swapping {
		swapSet[idx] = true
	}
	cmpSet := make(map[int]bool, len(step.Comparing))
	for _, idx := range step.Comparing {
		cmpSet[idx] = true
	}
	sortedSet := make(map[int]bool, len(step.Sorted))
	for _, idx := range step.Sorted {
		sortedSet[idx] = true
	}

	// Determine column width: 2 chars per column (consistent for alignment)
	colWidth := 2

	maxHeight := min(max(m.height-14, 6), 25)
	var lines []string

	for h := maxHeight; h > 0; h-- {
		var line strings.Builder
		for i, val := range step.Array {
			barHeight := int(float64(val) / float64(maxVal) * float64(maxHeight))

			char := " "
			if barHeight >= h {
				char = "█"
			}

			style := barStyle
			if swapSet[i] {
				style = swappingBarStyle
			} else if cmpSet[i] {
				style = comparingBarStyle
			} else if step.Pivot == i {
				style = pivotBarStyle
			} else if sortedSet[i] {
				style = sortedBarStyle
			}

			line.WriteString(style.Render(char))
			if colWidth > 1 {
				line.WriteString(" ")
			}
		}
		lines = append(lines, line.String())
	}

	if m.height >= 20 {
		var valLine strings.Builder
		for i, val := range step.Array {
			style := barStyle
			if sortedSet[i] {
				style = sortedBarStyle
			}
			label := fmt.Sprintf("%-*d", colWidth, val)
			valLine.WriteString(style.Render(label))
		}
		lines = append(lines, valLine.String())
	}

	return strings.Join(lines, "\n")
}

func (m Model) renderSortingLegend() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		barStyle.Render("█")+" Unsorted  ",
		comparingBarStyle.Render("█")+" Comparing  ",
		swappingBarStyle.Render("█")+" Swapping  ",
		pivotBarStyle.Render("█")+" Pivot  ",
		sortedBarStyle.Render("█")+" Sorted",
	)
}

func (m Model) viewPathfindingVisualization() string {
	if m.pathfindingAlgo == nil || len(m.pathfindingAlgo.Steps) == 0 {
		return "Loading..."
	}

	step := m.pathfindingAlgo.Steps[m.pathfindingStep]
	var b strings.Builder

	// Header (compact for small terminals)
	if m.height >= 25 {
		header := lipgloss.JoinHorizontal(lipgloss.Top,
			algoNameStyle.Render(m.pathfindingAlgo.Name),
			"  ",
			complexityStyle.Render("Time: "+m.pathfindingAlgo.TimeComplex),
		)
		b.WriteString(header)
		b.WriteString("\n")
		b.WriteString(algoDescStyle.Render(m.pathfindingAlgo.Description))
		b.WriteString("\n")
	} else {
		b.WriteString(algoNameStyle.Render(m.pathfindingAlgo.Name))
		b.WriteString("\n")
	}

	// Grid visualization
	vis := m.renderPathfindingGrid(step)
	b.WriteString(compactPanelStyle.Render(vis))
	b.WriteString("\n")

	// Status with frontier count
	progress := fmt.Sprintf("Step %d/%d", m.pathfindingStep+1, len(m.pathfindingAlgo.Steps))
	visitedCount := len(step.Visited)
	frontierCount := len(step.Frontier)
	statusLine := lipgloss.JoinHorizontal(lipgloss.Top,
		statusStyle.Render(progress),
		"  ",
		speedStyle.Render(m.speed.String()),
		"  ",
		lipgloss.NewStyle().Foreground(neonPurple).Render(fmt.Sprintf("Visited: %d", visitedCount)),
		"  ",
		lipgloss.NewStyle().Foreground(neonOrange).Render(fmt.Sprintf("Frontier: %d", frontierCount)),
		"  ",
		m.getPauseStatus(),
	)
	b.WriteString(statusLine)

	// Description (skip for small terminals)
	if m.height >= 22 {
		b.WriteString("\n")
		b.WriteString(descriptionStyle.Render("▸ " + step.Description))
	}
	b.WriteString("\n")

	// Legend (only for larger terminals)
	if m.height >= 25 {
		legend := m.renderPathfindingLegend()
		b.WriteString(legend)
		b.WriteString("\n")
	}

	controls := controlStyle.Render(
		keyStyle.Render("1/2/3") + " Speed  " +
			keyStyle.Render("p") + " Pause  " +
			keyStyle.Render("←/→") + " Step  " +
			keyStyle.Render("r") + " New Maze  " +
			keyStyle.Render("Esc") + " Back",
	)
	b.WriteString(controls)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, b.String())
}

func (m Model) renderPathfindingGrid(step pathfinding.Step) string {
	grid := step.Grid
	var lines []string

	frontierSet := make(map[pathfinding.Cell]bool, len(step.Frontier))
	for _, c := range step.Frontier {
		frontierSet[c] = true
	}

	pathSet := make(map[pathfinding.Cell]bool, len(step.Path))
	for _, c := range step.Path {
		pathSet[c] = true
	}

	for row := 0; row < grid.Height; row++ {
		var line strings.Builder
		for col := 0; col < grid.Width; col++ {
			cell := pathfinding.Cell{Row: row, Col: col}
			char := "·"
			style := emptyStyle

			if grid.Walls[cell] {
				char = "█"
				style = wallStyle
			} else if cell == grid.Start {
				char = "S"
				style = startStyle
			} else if cell == grid.End {
				char = "E"
				style = endStyle
			} else if pathSet[cell] {
				char = "●"
				style = pathStyle
			} else if cell == step.Current {
				char = "◆"
				style = currentStyle
			} else if frontierSet[cell] {
				char = "◇"
				style = frontierStyle
			} else if step.Visited[cell] {
				char = "○"
				style = visitedStyle
			}

			line.WriteString(style.Render(char))
			line.WriteString(" ")
		}
		lines = append(lines, line.String())
	}

	return strings.Join(lines, "\n")
}

func (m Model) renderPathfindingLegend() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		startStyle.Render("S")+" Start  ",
		endStyle.Render("E")+" End  ",
		wallStyle.Render("█")+" Wall  ",
		frontierStyle.Render("◇")+" Frontier  ",
		visitedStyle.Render("○")+" Visited  ",
		currentStyle.Render("◆")+" Current  ",
		pathStyle.Render("●")+" Path",
	)
}

func (m Model) getPauseStatus() string {
	isFinished := false
	switch m.view {
	case viewSortingVis:
		isFinished = m.sortingAlgo != nil && m.sortingStep >= len(m.sortingAlgo.Steps)-1
	case viewPathfindingVis:
		isFinished = m.pathfindingAlgo != nil && m.pathfindingStep >= len(m.pathfindingAlgo.Steps)-1
	}

	if isFinished {
		return lipgloss.NewStyle().Foreground(neonCyan).Bold(true).Render("✓ Complete")
	}
	if m.paused {
		return lipgloss.NewStyle().Foreground(neonOrange).Bold(true).Render("⏸ PAUSED")
	}
	return lipgloss.NewStyle().Foreground(neonGreen).Render("▶ Running")
}
