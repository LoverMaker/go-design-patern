package simple_factory

import "fmt"

type Animal interface {
	Say(name string) string
}

func NewAnimal(t int) (Animal){
	switch t {
	case 1: return &Bird{}
	case 2: return &Dog{}
	}
	return nil
}

type Bird struct {

}

func (*Bird)Say(name string) (string) {
	return fmt.Sprintf("bird")
}

type Dog struct {

}

func (*Dog)Say(name string) (string) {
	return fmt.Sprintf("dog")
}