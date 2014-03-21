package fields

import (
	"github.com/kirves/revel-forms/widgets"
)

type SubmitButtonType struct {
	Field
}

type ResetButtonType struct {
	Field
}

func SubmitButton(name string, text string) *SubmitButtonType {
	ret := &SubmitButtonType{
		FieldWithType(name, SUBMIT),
	}
	ret.SetText(text)
	return ret
}

func (b *SubmitButtonType) SetStyle(style string) {
	b.widget = widgets.Button(style)
}

func ResetButton(name string, text string) *ResetButtonType {
	ret := &ResetButtonType{
		FieldWithType(name, RESET),
	}
	ret.SetText(text)
	return ret
}

func (b *ResetButtonType) SetStyle(style string) {
	b.widget = widgets.Button(style)
}