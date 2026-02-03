package main

import "fmt"

type Robot struct {
	walk   Walkable
	talk   Talkable
	beahve Behaviour
}

type Walkable interface {
	walk() string
}

type NormalWalk struct{}

func (n *NormalWalk) walk() string {
	return "Walking normally"
}

type Nowalk struct{}

func (n *Nowalk) walk() string {
	return "Cannot walk"
}

type Talkable interface {
	talk() string
}

type NormalTalk struct{}

func (n *NormalTalk) talk() string {
	return "Talking normally"
}

type NoTalk struct{}

func (n *NoTalk) talk() string {
	return "Cannot talk"
}

type Behaviour interface {
	behave() string
}

type NormalBehaviour struct{}

func (n *NormalBehaviour) behave() string {
	return "Behaving normally"
}

type AggressiveBehaviour struct{}

func (a *AggressiveBehaviour) behave() string {
	return "Behaving aggressively"
}

func NewRobot(w Walkable, t Talkable, b Behaviour) Robot {
	return Robot{walk: w, talk: t, beahve: b}
}

func main() {
	robot1 := NewRobot(&NormalWalk{}, &NoTalk{}, &NormalBehaviour{})

	fmt.Println(robot1.walk.walk())
	fmt.Println(robot1.talk.talk())
	fmt.Println(robot1.beahve.behave())

	robot2 := NewRobot(&Nowalk{}, &NormalTalk{}, &AggressiveBehaviour{})

	fmt.Println(robot2.walk.walk())
	fmt.Println(robot2.talk.talk())
	fmt.Println(robot2.beahve.behave())
}
