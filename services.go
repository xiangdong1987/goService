package main

import (
	"errors"
	"goService/model"
)

type FilterService interface {
	Match(string) (string, string, error)
	Count(string) int
}
type filterService struct {
	AcTree *model.AcTrie
}

func (sv filterService) Match(s string) (string, string, error) {
	if s == "" {
		return "", "", ErrEmpty
	}
	//fmt.Println(sv.AcTree.Dictionary)
	level, str := sv.AcTree.Match(s)
	return level, str, nil
}

func (filterService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
