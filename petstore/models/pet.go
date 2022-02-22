// This package exposes functions to manage
// pets for petstore loneliness 2000.

package models

import (
	"fmt"
	"petstore/manage/utils"
	"reflect"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/jedib0t/go-pretty/table"
)

type Pet struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Animal      string    `json:"animal" db:"animal"`
	Breed       string    `json:"breed" db:"breed"`
	Age         int       `json:"age" db:"age"`
	TimeMeasure string    `json:"time_measure" db:"time_measure"`
	Price       float64   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Pets []Pet

const (
	GREATER_THAN PriceOperation = ">"
	LESS_THAN    PriceOperation = "<"
	EQUAL_TO     PriceOperation = "="
)

type PriceOperation string

func (p Pet) Create(animal, breed string, price float64, age int) {

}

const notDisplayedColumns = "CreatedAt, UpdatedAt, ID"

func (p Pets) Init(tx *pop.Connection) Pets {
	pets := Pets{
		{Animal: "Perro", Age: 12, TimeMeasure: "mm", Breed: "Pincher", Price: 100.23},
		{Animal: "Perro", Age: 6, TimeMeasure: "yy", Breed: "Doberman", Price: 250.23},
		{Animal: "Perro", Age: 4, TimeMeasure: "mm", Breed: "Husky", Price: 450.23},
		{Animal: "Perro", Age: 2, TimeMeasure: "mm", Breed: "Labrador", Price: 60.23},
		{Animal: "Perro", Age: 3, TimeMeasure: "mm", Breed: "Golden", Price: 80.23},
	}

	err := tx.Create(&pets)
	if err != nil {
		fmt.Println("err", err.Error())
		return Pets{}
	}

	return pets
}

func (p Pets) GenerateRows() []table.Row {
	rows := []table.Row{}

	for _, pet := range p {
		rows = append(rows, table.Row{pet.Animal, pet.Breed, pet.Age, pet.Price})
	}

	return rows
}

func (p Pets) NotDisplayedColumns() string {
	return notDisplayedColumns
}

func (p Pets) Display(style table.Style) {
	columns := Pet{}.GenerateColumns()
	rows := p.GenerateRows()

	utils.DisplayStruct(columns, rows, style)
}

func (p Pet) GenerateColumns() table.Row {
	s := reflect.ValueOf(&p).Elem()

	return utils.GenerateColumns(s, notDisplayedColumns)
}

func (p Pet) FindByAnimalName(name string, tx *pop.Connection) Pets {
	pets := Pets{}

	if err := tx.Where("animal = ?", name).All(&pets); err != nil {
		return Pets{}
	}

	return pets
}

func (p Pet) FindByPrice(operation PriceOperation, price float64, tx *pop.Connection) Pets {
	pets := Pets{}

	if err := tx.Where("price "+string(operation)+" ?", price).All(&pets); err != nil {
		return Pets{}
	}

	return pets
}

func (p Pet) FindByAge(age int, timeMeasure string, tx *pop.Connection) Pets {
	pets := Pets{}

	if err := tx.Where("age = ? AND time_measure = ?", age, timeMeasure).All(&pets); err != nil {
		return Pets{}
	}

	return pets
}
