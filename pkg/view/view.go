package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	selectedStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.NormalBorder()).Padding(0, 1).Width(10).BorderForeground(lipgloss.Color("#E83535"))
	selectStyle   = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).BorderStyle(lipgloss.NormalBorder()).Padding(0, 1).Width(10)
)

const (
	pageMain = iota
	pageHelp
	pageOptions
	pageProcess
)

const (
	modeFast = iota
	modeDebug
	modeLog
)

var ktransVersionList = []string{"v0.2.2", "v0.2.2rc1", "v0.2.2rc2", "v0.2.3", "v0.2.3post1", "v0.2.3post2", "v0.2.4", "v0.2.4post1"}

type Options struct {
	mode int

	installPath    string
	envInstallPath string
	condaBasePath  string
	envName        string

	maxJobs int
	useNuma bool

	useGhproxy bool
	ghproxyUrl string

	ktransVersion string

	logFile string
}

// 定义模型结构体
type model struct {
	page         int
	pageMainSate int
	screenWidth  int
	screenHeight int
	options      Options
}

// 初始化模型
func InitialModel() model {
	return model{
		page: pageMain,
		options: Options{
			mode: modeFast,

			installPath:    "",
			envInstallPath: "",
			condaBasePath:  "",
			envName:        "",

			maxJobs: 1,
			useNuma: false,

			useGhproxy: false,
			ghproxyUrl: "https://ghfast.top",

			ktransVersion: ktransVersionList[len(ktransVersionList)-1],

			logFile: "ktrans.log",
		},
	}
}

// 实现 Init 方法（可选）
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) UpdateMain(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.Type == tea.KeyEnter {
		switch m.pageMainSate {
		case 0:
			m.page = pageProcess
		case 1:
			m.page = pageHelp
		case 2:
			m.page = pageOptions
		case 3:
			return m, tea.Quit
		}
	}
	if msg.String() == "w" {
		if m.pageMainSate > 1 {
			m.pageMainSate -= 2
		}
	}
	if msg.String() == "a" {
		if m.pageMainSate > 0 {
			m.pageMainSate--
		}
	}
	if msg.String() == "s" {
		if m.pageMainSate < 2 {
			m.pageMainSate += 2
		}
	}
	if msg.String() == "d" {
		if m.pageMainSate < 3 {
			m.pageMainSate++
		}
	}
	return m, nil
}

func (m model) UpdateOptions(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m model) UpdateHelp(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) UpdateProcess(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

// 处理消息（按键/事件）
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		} else {
			switch m.page {
			case pageMain:
				return m.UpdateMain(msg)
			case pageOptions:
				return m.UpdateOptions(msg)
			case pageHelp:
				return m.UpdateHelp(msg)
			case pageProcess:
				return m.UpdateProcess(msg)
			}
		}
	}
	return m, nil
}

const logo = "╦╔═┌┬┐┬─┐┌─┐┌┐┌┌─┐┌─┐┌─┐┬─┐┌┬┐┌─┐┬─┐┌─┐\n" +
	"╠╩╗ │ ├┬┘├─┤│││└─┐├┤ │ │├┬┘│││├┤ ├┬┘└─┐\n" +
	"╩ ╩ ┴ ┴└─┴ ┴┘└┘└─┘└  └─┘┴└─┴ ┴└─┘┴└─└─┘"

// 定义视图
func (m model) pageMainView() string {
	btStyle := []lipgloss.Style{selectStyle, selectStyle, selectStyle, selectStyle}
	btStyle[m.pageMainSate] = selectedStyle
	res := lipgloss.JoinHorizontal(lipgloss.Center, btStyle[0].Render("run"), btStyle[1].Render("option"), btStyle[2].Render("help"), btStyle[3].Render("exit"))
	res = lipgloss.JoinVertical(lipgloss.Center, lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Padding(1).Render(logo), res)
	return lipgloss.Place(m.screenWidth, m.screenHeight, lipgloss.Center, lipgloss.Center, res)
}

func (m model) pageHelpView() string {
	return "help"
}

func (m model) pageOptionsView() string {
	return "options"
}

func (m model) pageProcessView() string {
	return "process"
}

// 渲染界面
func (m model) View() string {
	switch m.page {
	case pageHelp:
		return m.pageHelpView()
	case pageOptions:
		return m.pageOptionsView()
	case pageProcess:
		return m.pageProcessView()
	default:
		return m.pageMainView()
	}
}
