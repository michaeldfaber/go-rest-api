package seed

import (
	"log"

	"github.com/jinzhu/gorm"

	"go-rest-api/api/models"
)

var Persons = []*models.Person {
	&models.Person {
		Gender: "male",
		FirstName: "Michael",
		LastName: "Faber",
		Age: 24,
	},
	&models.Person {
		Gender: "male",
		FirstName: "Bob",
		LastName: "Johnson",
		Age: 30,
	},
	&models.Person {
		Gender: "female",
		FirstName: "Jane",
		LastName: "Smith",
		Age: 45,
	},
}

func Load(db *gorm.DB) {

	seeded := db.Debug().HasTable(&models.Person{})

	if seeded == false {
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
}