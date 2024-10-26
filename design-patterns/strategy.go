package designpatterns

import (
	"fmt"
	"strings"
)

// ohboyohboy

func TestStrategyPattern() {
	fmt.Println("Strategy Pattern")

	// create a new text process based on a custom strategy
	tp := NewTextProcessor(&HTMLListStrategy{})

	// adds items
	tp.AppendListItems([]string{"item 1", "item 2", "item 3"})

	// print result
	// prints a html <ul> <li> </li> </ul> element
	fmt.Println("result:\n", tp.String())

}

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

// for creating markdown lists
type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {
}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {
}

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf("* %s\n", item))
}

// for creating html lists
type HTMLListStrategy struct{}

func (m *HTMLListStrategy) Start(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("<ul>\n"))
}

func (m *HTMLListStrategy) End(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("</ul>\n"))
}

func (m *HTMLListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf("\t<li>%s</li>\n", item))
}

// for creating text processors with a chosen list strategy
type TextProcessor struct {
	builder      *strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{
		builder:      &strings.Builder{},
		listStrategy: listStrategy,
	}
}

type OutputFormat int

const (
	html OutputFormat = iota
	markdown
)

func (tp *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case html:
		tp.listStrategy = &HTMLListStrategy{}
	case markdown:
		tp.listStrategy = &MarkdownListStrategy{}
	}
}

// create the list regardless of if using HTML or Markdown strategy
func (tp *TextProcessor) AppendListItems(items []string) {
	tp.listStrategy.Start(tp.builder)

	for _, item := range items {
		tp.listStrategy.AddListItem(tp.builder, item)
	}

	tp.listStrategy.End(tp.builder)
}

// returns the string that has been built up
func (tp *TextProcessor) String() string {
	return tp.builder.String()
}
