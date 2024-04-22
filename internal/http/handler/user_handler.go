package handler

import (
	"atmail/internal/helper"
	"atmail/internal/model"
	"atmail/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const SUCCESS = "Successfully deleted"

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{
		userService: service,
	}
}

// @Summary 	Create User
// @Description Create User
// @Tags 		Users
// @Id 			Create
// @Produce 	json
// @Param 		Body  body  model.UserRequest  true  "User Details"
// @Router 		/users [post]
// @Success 	201 {object} model.User
// @Failure      400 {object} model.Error
// @Security 	BasicAuth
func (u *UserHandler) Create(ctx *gin.Context) {
	log.Infoln("Creating user...")
	var req model.UserRequest
	ctx.BindJSON(&req)
	if err := u.userService.ValidateNewUser(req); err != nil {
		log.Debugf("Validation failed: %+v %+v", err.Error(), req)
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	newUser, err := u.userService.Save(req)
	if err != nil {
		log.Debugf("Error creating user: %+v %+v", err.Error(), req)
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	log.Infoln("Successfully created user.")
	ctx.JSON(http.StatusCreated, newUser)
}

// @Summary      Retrieve user details by ID
// @Description  Retrieve user details by ID
// @Tags         Users
// @Id           Get
// @Produce      json
// @Param        id  path  string true "User ID"
// @Router       /users/{id} [get]
// @Success      200 {object} model.User
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Security BasicAuth
func (u *UserHandler) Get(ctx *gin.Context) {
	log.Infoln("Retrieving user details...")
	id, err := helper.CleanID(ctx.Param("id"))
	if err != nil {
		log.Debugf("Validation failed: %+v %+v", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	user, statusCode, err := u.userService.Get(*id)
	if err != nil {
		log.Debugf("Error retrieving user: %+v %+v", err.Error(), id)
		ctx.JSON(statusCode, model.Error{Error: err.Error()})
		return
	}
	log.Infoln("Done retrieving user details.")
	ctx.JSON(statusCode, user)
}

// @Summary      Retrieve all users
// @Description  Retrieve all users
// @Tags         Users
// @Id           GetAll
// @Produce      json
// @Router       /users [get]
// @Success      200 {object} model.User
// @Failure      400 {object} model.Error
// @Security BasicAuth
func (u *UserHandler) GetAll(ctx *gin.Context) {
	log.Infoln("Retrieving all users...")
	users, err := u.userService.GetAll()
	if err != nil {
		log.Debugf("Error retrieving user: %+v", err.Error())
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	log.Infoln("Done retrieving all users.")
	ctx.JSON(http.StatusOK, users)
}

// @Summary      Update User Dettails
// @Description  Update User Dettails
// @Tags         Users
// @Id           Update
// @Produce      json
// @Param        Body  body  model.UserRequest  true  "Update User"
// @Param        id  path  string true "User ID"
// @Router       /users/{id} [put]
// @Success      200 {object} model.User
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Security BasicAuth
func (u *UserHandler) Update(ctx *gin.Context) {
	log.Infoln("Updating user details...")
	id, err := helper.CleanID(ctx.Param("id"))
	if err != nil {
		log.Debugf("Validation failed: %+v %+v", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	var req model.User
	ctx.BindJSON(&req)
	req.ID = *id
	statusCode, err := u.userService.ValidateExistingUser(req)
	if err != nil {
		log.Debugf("Validation failed: %+v %+v", err.Error(), req)
		ctx.JSON(statusCode, model.Error{Error: err.Error()})
		return
	}

	newUser, err := u.userService.Update(req)
	if err != nil {
		log.Debugf("Error updating user: %+v %+v", err.Error(), req)
		ctx.JSON(statusCode, model.Error{Error: err.Error()})
		return
	}
	log.Infoln("Successfully updated user details.")
	ctx.JSON(http.StatusOK, newUser)
}

// @Title        Delete User
// @Summary      Delete User
// @Description  Delete User
// @Tags         Users
// @Id           Delete
// @Produce      json
// @Param        id  path  string true "User ID"
// @Router       /users/{id} [delete]
// @Success      200 string string
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Security BasicAuth
func (u *UserHandler) Delete(ctx *gin.Context) {
	log.Infoln("Deleting user...")
	id, err := helper.CleanID(ctx.Param("id"))
	if err != nil {
		log.Debugf("Validation failed: %+v %+v", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	statusCode, err := u.userService.ValidateID(*id)
	if err != nil {
		log.Debugf("Validation failed: %+v %+v", err.Error(), id)
		ctx.JSON(statusCode, model.Error{Error: err.Error()})
		return
	}

	if err := u.userService.Delete(*id); err != nil {
		log.Debugf("Error deleting user: %+v %+v", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	log.Infoln("Successfully deleted user...")
	ctx.JSON(http.StatusOK, SUCCESS)
}
