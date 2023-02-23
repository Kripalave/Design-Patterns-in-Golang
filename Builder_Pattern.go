package main
import "fmt"

type MyAudi struct {
	name string
	wheelSize string
	engineCapacityCC int
	powerBHP int
	torque int
	fuelType string
	cost int
}

func (m MyAudi) showAudi() {
	fmt.Println("Audi Name : " , m.name)
	fmt.Println("Audi wheel : " , m.wheelSize)
	fmt.Println("Audi Engine : " , m.engineCapacityCC)
	fmt.Println("Audi power bhp : " , m.powerBHP)
	fmt.Println("Audi Torque : " , m.torque)
	fmt.Println("Audi fuel : " , m.fuelType)
	fmt.Println("Audi Cost INR: " , m.cost)

}

//this is abstract builder
type AudiBuilder interface {
	createAudi() MyAudi
}

//now concrete builders
type AudiQ3Builder struct {
}

func (a AudiQ3Builder) createAudi() MyAudi {
	audi := MyAudi{}
	audi.name = "AydiQ3"
	audi.cost = 1000000000
	audi.engineCapacityCC = 3000
	audi.fuelType = "Diesel"
	audi.powerBHP = 300
	audi.torque = 480
	audi.wheelSize = "R17/75/250"
	return audi
}


//now concrete builders
type AudiQ4Builder struct {
}

func (a AudiQ4Builder) createAudi() MyAudi {
	audi := MyAudi{}
	audi.name = "AydiQ4"
	audi.cost = 2000000000
	audi.engineCapacityCC = 4000
	audi.fuelType = "Patrol"
	audi.powerBHP = 400
	audi.torque = 580
	audi.wheelSize = "R19/85/250"
	return audi
}


// you may implement other builder urself :P


//this is director

type ProductionLine struct {
	builder AudiBuilder
}

func (a ProductionLine)CreateMyAudi() MyAudi {
	return a.builder.createAudi()
}

func main() {
	p := ProductionLine{AudiQ3Builder{}}
	res := p.CreateMyAudi()
	res.showAudi()
	p = ProductionLine{AudiQ4Builder{}}
	res = p.CreateMyAudi()
	res.showAudi()

}