package entity

import "github.com/MsSabo/simulation/internal"

type Entity interface {
	internal.ICell
	GetSign() string
}