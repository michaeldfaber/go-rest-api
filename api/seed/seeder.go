package seed

import (
	"log"

	"github.com/jinzhu/gorm"

	"go-rest-api/api/models"
)

var Persons = []*models.Person {
	&models.Person { 
		Id: 1,
		Gender: "male",
		FirstName: "Michael",
		LastName: "Faber",
		Age: 24,
	},
	&models.Person { 
		Id: 2,
		Gender: "male",
		FirstName: "Bob",
		LastName: "Johnson",
		Age: 30,
	},
	&models.Person { 
		Id: 3,
		Gender: "female",
		FirstName: "Jane",
		LastName: "Smith",
		Age: 45,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Person{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Person{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	for i, _ := range Persons {
		err = db.Debug().Model(&models.Person{}).Create(&Persons[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed Persons table: %v", err)
		}
	}
}