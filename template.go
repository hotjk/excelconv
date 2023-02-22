package main

import "strconv"

type Position struct {
	Name   string `json:"name"`
	Row    int    `json:"row"`
	Column int    `json:"column"`
	Fix    bool   `json:"fix"`
	Value  string `json:"value"`
}

type Cell struct {
	Position
	Func   string        `json:"func"`
	Params []interface{} `json:"params"`
}

type Link struct {
	To   Cell   `json:"to"`
	From []Cell `json:"from"`
}

type Loop struct {
	Stop  Position `json:"stop"`
	Links []Link   `json:"links"`
}

type Sheet struct {
	From      string `json:"from"`
	To        string `json:"to"`
	FromIndex int    `json:"fromIndex"`
	ToIndex   int    `json:"toIndex"`
	Links     []Link `json:"links"`
	Loops     []Loop `json:"loops"`
}

func (p *Position) SetName() {
	if p.Name == "" {
		p.Name = Index2Name(p.Column) + strconv.Itoa(p.Row)
	} else {
		if p.Name[0] == '$' {
			p.Fix = true
			p.Name = p.Name[1:]
		}
		p.Column, p.Row, _ = NameSplit(p.Name)
	}
}

func (link *Link) SetName() {
	link.To.SetName()
	for j := 0; j < len(link.From); j++ {
		from := &link.From[j]
		from.SetName()
	}
}

func (loop *Loop) SetName() {
	loop.Stop.SetName()
	for i := 0; i < len(loop.Links); i++ {
		link := &loop.Links[i]
		link.SetName()
	}
}

func (sheet *Sheet) SetName() {
	for i := 0; i < len(sheet.Links); i++ {
		link := &sheet.Links[i]
		link.SetName()
	}
	for i := 0; i < len(sheet.Loops); i++ {
		loop := &sheet.Loops[i]
		loop.SetName()
	}
}
