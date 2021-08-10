package models

import (
	"time"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Person struct {
	Id 			uint32 		`gorm:"primary_key;auto_increment" json:"Id"`
	Gender 		string 		`gorm:"size:6;not null;" json:"Gender"`
	FirstName 	string 		`gorm:"size:100;not null;" json:"FirstName"`
	LastName 	string 		`gorm:"size:100;not null;" json:"LastName"`
	Age 		int 		`gorm:"size:3;not null;" json:"Age"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"CreatedAt"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"UpdatedAt"`
}

func (p *Person) Prepare() {
	p.Id = 0
	p.Gender = html.EscapeString(strings.TrimSpace(p.Gender))
	p.FirstName = html.EscapeString(strings.TrimSpace(p.FirstName))
	p.LastName = html.EscapeString(strings.TrimSpace(p.LastName))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Person) SavePerson(db *gorm.DB) (*Person, error) {

	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Person{}, err
	}
	return p, nil
}

func (p *Person) FindAllPersons(db *gorm.DB) (*[]Person, error) {
	var err error
	persons := []Person{}
	err = db.Debug().Model(&Person{}).Limit(100).Find(&persons).Error
	if err != nil {
		return &[]Person{}, err
	}
	return &persons, err
}

// func (p *Person) Validate(action string) error {
// 	switch strings.ToLower(action) {
// 	case "create":
// 		if u.FirstName == "" {
// 			return errors.New("Required FirstName")
// 		}
// 		return nil
// 	case "update":
// 		return nil
// 	default:
// 		return nil
// 	}
// }

