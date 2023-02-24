package main

import "fmt"

type Mobile interface {
	setCamera(Camera)
	click()
}

type Samsung struct {
	Camera Camera
}

func (s *Samsung) setCamera(c Camera) {
	s.Camera = c
}

func (s *Samsung) click() {
	fmt.Println("Clicking via samsung phone")
	s.Camera.clickPic()
}

type Motorola struct {
	Camera Camera
}

func (m *Motorola) setCamera(c Camera) {
	m.Camera = c
}

func (m *Motorola) click() {
	fmt.Println("CLicking via motorola phone")
	m.Camera.clickPic()
}

type Camera interface {
	clickPic()
}

type FrontCamera struct {
}

func (f *FrontCamera) clickPic() {
	fmt.Println("Picture clicked via front camera")
}

type BackCamera struct {
}

func (f *BackCamera) clickPic() {
	fmt.Println("Picture clicked via back camera")
}
func main() {
	s := new(Samsung)
	m := new(Motorola)
	fc := new(FrontCamera)
	b := new(BackCamera)

	//samsung with front camera
	s.setCamera(fc)
	s.click()
	fmt.Println()
	//samsung with back camera
	s.setCamera(b)
	s.click()
	fmt.Println()
	//motorola with fron camera
	m.setCamera(fc)
	m.click()
	fmt.Println()
	//motorola with back camera
	m.setCamera(b)
	m.click()


}