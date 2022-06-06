package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record Found")

type Tarefas struct{
  ID      int
  Title   string
  Content string
}