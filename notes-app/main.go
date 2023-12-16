package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/robertjbass/notes/note"
	"github.com/robertjbass/notes/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed")
		return err
	}

	fmt.Println("Saving succeeded")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title")
	content := getUserInput("Note content")

	return title, content
}

func main() {
	todoText := getUserInput("Todo")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		return
	}

	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func getUserInput(prompt string) string {
	fmt.Printf("%s: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
