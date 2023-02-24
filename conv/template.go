package conv

import "strconv"

type position struct {
	Name  string `json:"name"`
	row   int
	col   int
	fix   bool
	Value string `json:"value"`
}

type cell struct {
	position
	Func   string        `json:"func"`
	Params []interface{} `json:"params"`
}

type link struct {
	To   cell   `json:"to"`
	From []cell `json:"from"`
}

type loop struct {
	Stop  position `json:"stop"`
	Links []link   `json:"links"`
}

type sheet struct {
	From      string
	To        string
	FromIndex int    `json:"fromIndex"`
	ToIndex   int    `json:"toIndex"`
	Links     []link `json:"links"`
	Loops     []loop `json:"loops"`
}

func (p *position) setName() {
	if p.Name == "" {
		p.Name = Index2Name(p.col) + strconv.Itoa(p.row)
	} else {
		if p.Name[0] == '$' {
			p.fix = true
			p.Name = p.Name[1:]
		}
		p.col, p.row, _ = NameSplit(p.Name)
	}
}

func (l *link) setName() {
	l.To.setName()
	for j := 0; j < len(l.From); j++ {
		from := &l.From[j]
		from.setName()
	}
}

func (l *loop) setName() {
	l.Stop.setName()
	for i := 0; i < len(l.Links); i++ {
		link := &l.Links[i]
		link.setName()
	}
}

func (s *sheet) SetName() {
	for i := 0; i < len(s.Links); i++ {
		link := &s.Links[i]
		link.setName()
	}
	for i := 0; i < len(s.Loops); i++ {
		loop := &s.Loops[i]
		loop.setName()
	}
}
