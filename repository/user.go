package repository

import (
	"database/sql"
	"log"
	"lunch-be/model"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{Db: db}
}

func (r *UserRepository) DeleteUser(id uint) bool {
	_, err := r.Db.Exec("DELETE FROM users WHERE user_id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *UserRepository) GetOneUser(id uint) model.User {
	query, err := r.Db.Query("SELECT * FROM users WHERE user_id = $1", id)
	if err != nil {
		log.Println(err)
		return model.User{}
	}
	var user model.User
	if query != nil {
		for query.Next() {
			var (
				id         uint
				voornaam   string
				achternaam string
			)
			err := query.Scan(&id, &voornaam, &achternaam)
			if err != nil {
				log.Println(err)
			}
			user = model.User{Id: id, VoorNaam: voornaam, FamilieNaam: achternaam}
		}
	}
	return user
}

func (r *UserRepository) GetAllUsers() []model.User {
	query, err := r.Db.Query("SELECT * FROM users ORDER BY voornaam")
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []model.User
	if query != nil {
		for query.Next() {
			var (
				id         uint
				voornaam   string
				achternaam string
			)
			err := query.Scan(&id, &voornaam, &achternaam)
			if err != nil {
				log.Println(err)
			}
			resto := model.User{Id: id, VoorNaam: voornaam, FamilieNaam: achternaam}
			users = append(users, resto)
		}
	}
	return users
}

func (r *UserRepository) InsertUser(post model.PostUser) bool {
	statement, err := r.Db.Prepare("INSERT INTO users(voornaam, achternaam) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer statement.Close()

	_, err2 := statement.Exec(post.VoorNaam, post.FamilieNaam)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

func (r *UserRepository) UpdateUser(id uint, post model.UpdateUser) model.User {
	_, err := r.Db.Exec("UPDATE users SET voornaam = $1, achternaam = $2 WHERE user_id = $3", post.VoorNaam, post.FamilieNaam, id)
	if err != nil {
		log.Println(err)
		return model.User{}
	}
	return r.GetOneUser(id)
}
