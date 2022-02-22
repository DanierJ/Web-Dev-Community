package client

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

		clients := []models.Client{
			{Name: "John", LastName: "Doe", Email: "johndoe@test.com", Phone: "123-777-8765", Gender: "M", Age: 24, Address: "St Jordi 234"},
		}

		err := tx.Create(&clients)
		if err != nil {
			return
		}

		var createdClients []models.Client
		r.NoError(tx.All(&createdClients))

		r.Equal(1, len(createdClients))
	})
}

func Test_List(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		clients := models.Clients{}.Init(tx)
		createdClients := models.Clients{}
		r.NoError(tx.All(&createdClients))

		r.Equal(len(createdClients), len(createdClients))

		for i, client := range clients {
			r.Equal(client.Name, createdClients[i].Name)
		}

		createdClients.Display(table.StyleLight)
	})
}

func Test_Find(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		clients := models.Clients{}.Init(tx)
		r.Equal(true, len(clients) > 0)

		clientsFoundByNameLastname := models.Client{}.FindByNameLastname("pepo", tx)
		clientsFoundByEmail := models.Client{}.FindByEmail("my email", tx)
		clientsPetsFound := models.Client{}.FindPets(tx)

		clientsFoundByNameLastname.Display(table.StyleLight)
		clientsFoundByEmail.Display(table.StyleLight)
		clientsPetsFound.Display(table.StyleLight)

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
		clients := models.Clients{}.Init(tx)

		r.Equal(true, len(clients) > 0)

		cID := clients[0].ID

		r.NoError(tx.Destroy(&clients[0]))

		c := models.Client{}

		r.Error(tx.Find(&c, cID))

		r.Empty(c)

		//r.Fail("Que se puedan remover o eliminar entidades")
	})
}

func Test_Update(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		clients := models.Clients{}.Init(tx)
		r.Equal(true, len(clients) > 0)

		oldClient := clients[0]

		clients[0].Name = "Robb"

		r.NoError(tx.Update(&clients[0]))

		r.NotEqual(oldClient.Name, clients[0].Name)

		//r.Fail("Que se puedan actualizar entidades")
	})
}
