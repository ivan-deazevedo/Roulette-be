package controller

import (
	"database/sql"
	"lunch-be/model"
	"lunch-be/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Db *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &UserController{Db: db}
}

func (r *UserController) DeleteUser(c *gin.Context) {
	DB := r.Db
	var uri model.UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repo := repository.NewUserRepository(DB)
	delete := repo.DeleteUser(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "deleted user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete user failed"})
		return
	}

}

func (r *UserController) GetAllUsers(c *gin.Context) {
	DB := r.Db
	repo := repository.NewUserRepository(DB)
	get := repo.GetAllUsers()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get users successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "users not found"})
		return
	}
}

func (r *UserController) GetOneUser(c *gin.Context) {
	DB := r.Db
	var uri model.UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err.Error()})
		return
	}
	repo := repository.NewUserRepository(DB)
	get := repo.GetOneUser(uri.ID)
	if (get != model.User{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get user successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "user not found"})
		return
	}
}

func (r *UserController) InsertUser(c *gin.Context) {
	DB := r.Db
	var post model.PostUser
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repo := repository.NewUserRepository(DB)
	insert := repo.InsertUser(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert user failed"})
		return
	}
}

func (r *UserController) UpdateUser(c *gin.Context) {
	DB := r.Db
	var post model.UpdateUser
	var uri model.UserUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewUserRepository(DB)
	update := repository.UpdateUser(uri.ID, post)
	if (update != model.User{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update user failed"})
		return
	}
}
