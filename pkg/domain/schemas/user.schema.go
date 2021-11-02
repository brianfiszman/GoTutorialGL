package schemas

import (
	"reflect"
)

type UserSchema struct {
	Username string `json:"username" db:"username"`
	Fullname string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email"`
}

const (
	DatabaseTag = "db"
)

func NewUserSchema(u UserSchema) (copy UserSchema) {
	v := reflect.ValueOf(u)

	for i := 0; i < v.NumField(); i++ {
		switch v.Type().Field(i).Name {
		case "Username":
			if u.Username != "" {
				copy.Username = u.Username
			} else {
				copy.Username = "NULL"
			}
		case "Fullname":
			if u.Fullname != "" {
				copy.Fullname = u.Fullname
			} else {
				copy.Fullname = "NULL"
			}
		case "Email":
			if u.Email != "" {
				copy.Email = u.Email
			} else {
				copy.Email = "NULL"
			}
		}
	}
	return u
}

func (u UserSchema) StructToMap() map[string]interface{} {

	m := make(map[string]interface{})

	elem := reflect.ValueOf(u)

	for i := 0; i < elem.NumField(); i++ {
		m[elem.Type().Field(i).Tag.Get(DatabaseTag)] = elem.Field(i).Interface()
	}
	return m
}

func (u UserSchema) GetFields() []string {
	v := reflect.ValueOf(u)
	var fields []string

	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, v.Type().Field(i).Tag.Get(DatabaseTag))
	}

	return fields
}

func (u UserSchema) GetValues() (values []interface{}) {
	v := reflect.ValueOf(u)

	for i := 0; i < v.NumField(); i++ {
		values = append(values, v.Field(i).Interface())
	}

	return
}
