package pet_test

import (
	"log"
	"petstore/models"
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/jedib0t/go-pretty/table"
	"github.com/stretchr/testify/require"
)

var db *pop.Connection

func init() {
	pop.Debug = true
	pop.AddLookupPaths("../../")

	if err := pop.LoadConfigFile(); err != nil {
		log.Panic(err)
	}

	var err error
	db, err = pop.Connect("test")
	if err != nil {
		log.Panic(err)
	}
}

func transaction(fn func(tx *pop.Connection)) {
	err := db.Rollback(func(tx *pop.Connection) {
		fn(tx)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Create(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)

		pets := []models.Pet{
			{Animal: "Perro", Age: 12, Breed: "Pincher", Price: 200.23},
		}

		err := tx.Create(&pets)
		if err != nil {
			return
		}

		var createdPets []models.Pet
		r.NoError(tx.All(&createdPets))

		r.Equal(1, len(createdPets))
	})
}

func Test_List(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		pets := models.Pets{}.Init(tx)
		createdPets := models.Pets{}
		r.NoError(tx.All(&createdPets))

		r.Equal(len(pets), len(createdPets))

		for i, pet := range pets {
			r.Equal(pet.Animal, createdPets[i].Animal)
		}

		createdPets.Display(table.StyleLight)
	})
}

func Test_Find(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		pets := models.Pets{}.Init(tx)
		r.Equal(true, len(pets) > 0)

		petsFoundByAnimalName := models.Pet{}.FindByAnimalName("Perro", tx)
		petsFoundByPrice := models.Pet{}.FindByPrice(models.GREATER_THAN, 200.3, tx)
		petsFoundByAge := models.Pet{}.FindByAge(12, "mm", tx)

		petsFoundByAnimalName.Display(table.StyleLight)
		petsFoundByPrice.Display(table.StyleLight)
		petsFoundByAge.Display(table.StyleLight)

		// r.Fail(`Animal. Ejemplo: mostrar las mascotas que sean 'perros' o mostrar las que sean 'gatos'
		// Precio (Hasta X monto inclusive). Ejemplo :mostrar las mascotas que estén por debajo de $700.000.
		// Edad (meses o años): mostrar las mascotas que tengan X meses o Y años (es responsabilidad del usuario especificar si es mes o año)
		// `)
	})
}

func Test_Destroy(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		pets := models.Pets{}.Init(tx)

		r.Equal(true, len(pets) > 0)

		pID := pets[0].ID

		r.NoError(tx.Destroy(&pets[0]))

		p := models.Pet{}

		r.Error(tx.Find(&p, pID))

		r.Empty(p)

		//r.Fail("Que se puedan remover o eliminar entidades")
	})
}

func Test_Update(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		pets := models.Pets{}.Init(tx)
		r.Equal(true, len(pets) > 0)

		oldPet := pets[0]

		pets[0].Price = 700.5

		r.NoError(tx.Update(&pets[0]))

		r.NotEqual(oldPet.Price, pets[0].Price)

		//r.Fail("Que se puedan actualizar entidades")
	})
}
