package main

import (
	"fmt"
)

type Animal interface {
	Sound(typer string) string
	Nutrition(feed string) string
}

type Human struct {
	Race string
	Feed string
}

type Dog struct {
	Breed string
	Feed  string
}

func (h *Human) Sound(typer string) string {
	return fmt.Sprintf("humans do %s", typer)
}

func (h *Human) Nutrition(feed string) string {
	return fmt.Sprintf("humans can be %s", feed)
}

func (d *Dog) Sound(typer string) string {
	return fmt.Sprintf("Dogs do %s", typer)
}

func (d *Dog) Nutrition(feed string) string {
	return fmt.Sprintf("Dogs are %s", feed)
}

func AnimalSound(animal Animal, sound string) {
	fmt.Println(animal.Sound(sound))
}

func AnimalFeed(animal Animal, feed string){
	fmt.Println(animal.Nutrition(feed))
}
