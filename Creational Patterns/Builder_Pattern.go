package main

import "fmt"

// ReportBuilder is the builder interface for constructing reports.
type ReportBuilder interface {
	SetTitle(title string) ReportBuilder
	SetContent(content string) ReportBuilder
	Build() string
}

// ConcreteReportBuilder is a concrete builder implementation for reports.
type ConcreteReportBuilder struct {
	title   string
	content string
}

// NewReportBuilder creates a new instance of ConcreteReportBuilder.
func NewReportBuilder() ReportBuilder {
	return &ConcreteReportBuilder{}
}

func (b *ConcreteReportBuilder) SetTitle(title string) ReportBuilder {
	b.title = title
	return b
}

func (b *ConcreteReportBuilder) SetContent(content string) ReportBuilder {
	b.content = content
	return b
}

func (b *ConcreteReportBuilder) Build() string {
	return fmt.Sprintf("Title: %s\nContent: %s", b.title, b.content)
}

func main() {
	// Example usage of the report builder
	builder := NewReportBuilder().
		SetTitle("Monthly Report").
		SetContent("This is the content of the report.")

	report := builder.Build()
	fmt.Println(report)
}
