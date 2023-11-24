package main

import "fmt"

// GraphicObject represents a graphic object
type GraphicObject struct {
	Name  string
	Color string
	// Additional properties...
}

// Cloneable defines the prototype interface
type Cloneable interface {
	Clone() Cloneable
}

// NewGraphicObject creates a new instance of GraphicObject
func NewGraphicObject(name, color string) *GraphicObject {
	return &GraphicObject{Name: name, Color: color}
}

// Clone creates a shallow copy of the GraphicObject
func (g *GraphicObject) Clone() Cloneable {
	// Perform a shallow copy to create a new instance with the same values
	return &GraphicObject{Name: g.Name, Color: g.Color}
}

func main() {
	// Creating an instance of GraphicObject
	originalGraphic := NewGraphicObject("Circle", "Red")

	// Cloning the original graphic
	clonedGraphic := originalGraphic.Clone().(*GraphicObject)

	// Modifying the cloned graphic
	clonedGraphic.Color = "Blue"

	// Displaying the results
	fmt.Println("Original Graphic:", originalGraphic)
	fmt.Println("Cloned Graphic:", clonedGraphic)
}
