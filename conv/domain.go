package conv

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Domain struct {
	source   *excelize.File
	target   *excelize.File
	template *sheet
}

func New(sourceFile string, templateFile string) Domain {
	source, err := excelize.OpenFile(sourceFile)
	if err != nil {
		panic("open source file failed: " + sourceFile)
	}

	target := excelize.NewFile()

	dat, _ := os.ReadFile(templateFile)
	var template sheet
	json.Unmarshal(dat, &template)

	if template.From == "" {
		template.From = source.GetSheetName(int(template.FromIndex - 1))
	}
	if template.To == "" {
		template.To = target.GetSheetName(int(template.ToIndex - 1))
	}

	template.SetName()

	return Domain{source: source, target: target, template: &template}
}

func (d Domain) Convert() {
	for _, link := range d.template.Links {
		var value string
		for _, from := range link.From {
			v, _ := d.source.GetCellValue(d.template.From, from.Name)
			//fmt.Println("Link:", d.Template.From, from.Name, v)
			v = Invoke(v, from.Func, from.Params)
			value = value + v
		}
		value = Invoke(value, link.To.Func, link.To.Params)
		d.target.SetCellValue(d.template.To, link.To.Name, value)
	}

	for _, loop := range d.template.Loops {
		step := 0
		for {
			name := Index2Name(loop.Stop.Column) + strconv.Itoa(loop.Stop.Row+step)
			//fmt.Println(loop.Stop.Column, Index2Name(loop.Stop.Column), strconv.Itoa(loop.Stop.Row+step))
			if v, _ := d.source.GetCellValue(d.template.From, name); v == "" {
				break
			}
			for _, link := range loop.Links {
				var value string
				var v string
				for _, from := range link.From {
					if from.Value != "" {
						v = from.Value
					} else if from.Fix {
						v, _ = d.source.GetCellValue(d.template.From, from.Name)
					} else {
						name := Index2Name(from.Column) + strconv.Itoa(from.Row+step)
						v, _ = d.source.GetCellValue(d.template.From, name)
					}
					v = Invoke(v, from.Func, from.Params)
					//fmt.Println(v, from.Func, from.Params)
					value = value + v
				}
				name := Index2Name(link.To.Column) + strconv.Itoa(link.To.Row+step)
				//fmt.Println(name, value)
				value = Invoke(value, link.To.Func, link.To.Params)
				d.target.SetCellValue(d.template.To, name, value)
			}
			step++
		}
	}
}

func (d Domain) Save(filename string) {
	d.target.SaveAs(filename)
}

func (d Domain) Dispose() {
	d.source.Close()
	d.target.Close()
}
