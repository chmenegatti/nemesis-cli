package ui

import (
	"bytes"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

type App struct {
	// Contains the UI elements
	app        *tview.Application
	inputField *tview.InputField
	outputBox  *tview.TextView
	rootCmd    *cobra.Command
}

func NewApp() *App {
	a := &App{
		app:        tview.NewApplication(),
		inputField: tview.NewInputField(),
		outputBox:  tview.NewTextView(),
	}

	a.InitializeUI()
	return a
}

func (a *App) InitializeUI() {
	a.inputField.SetLabel("Nemesis: ").SetBorder(true).SetBorderColor(tcell.ColorBlue)
	a.outputBox.SetBorder(true)

	a.inputField.SetDoneFunc(
		func(key tcell.Key) {
			command := a.inputField.GetText()
			if key == tcell.KeyEnter {
				a.processCommand(command)
			} else if key == tcell.KeyCtrlC {
				a.app.Stop()
			}
			a.inputField.SetText("")
		},
	)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.inputField, 3, 1, true).
		AddItem(a.outputBox, 0, 1, false)

	a.app.SetRoot(flex, true).EnableMouse(true)
}

func (a *App) processCommand(command string) error {
	args := strings.Fields(command)
	if len(args) == 0 {
		return nil
	}

	buf := new(bytes.Buffer)
	a.rootCmd.SetOut(buf)
	a.rootCmd.SetErr(buf)
	a.rootCmd.SetArgs(args)

	err := a.rootCmd.Execute()
	if err != nil {
		a.outputBox.SetText("Erro: " + err.Error())
	} else {
		output := buf.String()
		if output == "" {
			output = "Comando executado sem saída."
		}
		a.outputBox.SetText(output)
	}
	return nil
}

func (a *App) SetRootCommand(cmd *cobra.Command) {
	a.rootCmd = cmd
}

func (a *App) Run() error {
	return a.app.Run()
}
