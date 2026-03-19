# Algo Arena

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Terminal-4D4D4D?style=for-the-badge&logo=gnome-terminal&logoColor=white" alt="Terminal">
  <img src="https://img.shields.io/badge/Bubble%20Tea-FF69B4?style=for-the-badge" alt="Bubble Tea">
  <img src="https://img.shields.io/badge/License-MIT-blue?style=for-the-badge" alt="License">
</p>

<p align="center">
  <b>Watch algorithms come to life in your terminal</b>
</p>

<p align="center">
  A beautiful, interactive CLI tool for visualizing sorting and pathfinding algorithms in real-time.
  <br>
  Built with Go and the <a href="https://github.com/charmbracelet/bubbletea">Bubble Tea</a> framework.
</p>

---

```
   █████╗ ██╗      ██████╗  ██████╗      █████╗ ██████╗ ███████╗███╗   ██╗ █████╗ 
  ██╔══██╗██║     ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██╔════╝████╗  ██║██╔══██╗
  ███████║██║     ██║  ███╗██║   ██║    ███████║██████╔╝█████╗  ██╔██╗ ██║███████║
  ██╔══██║██║     ██║   ██║██║   ██║    ██╔══██║██╔══██╗██╔══╝  ██║╚██╗██║██╔══██║
  ██║  ██║███████╗╚██████╔╝╚██████╔╝    ██║  ██║██║  ██║███████╗██║ ╚████║██║  ██║
  ╚═╝  ╚═╝╚══════╝ ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝
```

## Features

### Sorting Algorithms
Watch data transform in real-time with beautiful bar visualizations:

| Algorithm | Time Complexity | Space Complexity | Description |
|-----------|-----------------|------------------|-------------|
| **Bubble Sort** | O(n²) | O(1) | Repeatedly swaps adjacent elements |
| **Selection Sort** | O(n²) | O(1) | Finds minimum and places at beginning |
| **Insertion Sort** | O(n²) | O(1) | Builds sorted array one element at a time |
| **Quick Sort** | O(n log n) avg | O(log n) | Divide and conquer with pivot partitioning |
| **Merge Sort** | O(n log n) | O(n) | Recursive splitting and merging |
| **Heap Sort** | O(n log n) | O(1) | Uses binary heap data structure |

### Pathfinding Algorithms
Navigate through randomly generated mazes:

| Algorithm | Time Complexity | Description |
|-----------|-----------------|-------------|
| **BFS** | O(V + E) | Explores all neighbors at current depth first |
| **DFS** | O(V + E) | Explores as deep as possible before backtracking |
| **Dijkstra** | O((V + E) log V) | Finds shortest path by distance |
| **A*** | O(E log V) | Uses heuristic for optimal pathfinding |

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/ethanwritescode/algo-arena.git
cd algo-arena

# Install dependencies
go mod tidy

# Build
go build -o algo-arena

# Run
./algo-arena
```

### Using Go Install

```bash
go install github.com/ethanwritescode/algo-arena@latest
```

## Usage

Launch the application:

```bash
./algo-arena
```

### Controls

| Key | Action |
|-----|--------|
| `↑` / `↓` or `k` / `j` | Navigate menus |
| `Enter` or `Space` | Select option |
| `1` / `2` / `3` | Set speed (Slow/Normal/Fast) |
| `p` | Pause/Resume animation |
| `r` | Reset/Generate new data |
| `Esc` | Go back |
| `q` | Quit to menu / Exit |

## Visual Legend

### Sorting Visualization
- 🟦 **Cyan** - Unsorted elements
- 🟧 **Orange** - Elements being compared
- 🟪 **Pink** - Elements being swapped
- 🟨 **Yellow** - Pivot element (Quick Sort)
- 🟩 **Green** - Sorted elements

### Pathfinding Visualization
- **S** - Start position
- **E** - End/Goal position
- █ - Wall/Obstacle
- ○ - Visited cell
- ◆ - Current cell being explored
- ● - Final path

## Architecture

```
algo-arena/
├── main.go                          # Entry point
├── internal/
│   ├── algorithms/
│   │   ├── sorting/
│   │   │   └── sorting.go           # Sorting algorithms & step generation
│   │   └── pathfinding/
│   │       └── pathfinding.go       # Pathfinding algorithms & grid logic
│   └── tui/
│       ├── model.go                 # Bubble Tea model & update logic
│       └── styles.go                # Lipgloss styles & theming
├── go.mod
└── README.md
```

## Educational Value

This tool is perfect for:
- **CS Students** - Understand algorithm behavior visually
- **Interview Prep** - See how classic algorithms work
- **Teaching** - Demonstrate algorithms in class
- **Self-Learning** - Reinforce algorithmic concepts

## Contributing

Contributions are welcome! Here are some ideas:

- [ ] Add more sorting algorithms (Radix, Counting, Shell)
- [ ] Add more pathfinding algorithms (Greedy BFS, Jump Point Search)
- [ ] Add data structure visualizations (BST, Linked List, Hash Table)
- [ ] Add algorithm comparison mode
- [ ] Add step-by-step mode with explanations
- [ ] Export visualization as GIF

## License

MIT License - feel free to use this project for learning and portfolio purposes!

## Acknowledgments

- [Charm](https://charm.sh/) - For the amazing Bubble Tea and Lip Gloss libraries
- The Go community for excellent tooling

---

<p align="center">
  Made with 💜 by <a href="https://github.com/ethanwritescode">Ethan James</a>
</p>

<p align="center">
  ⭐ Star this repo if you found it helpful!
</p>
