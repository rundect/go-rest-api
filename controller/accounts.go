package controller

import (
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-rest-api/httputil"
	"go-rest-api/model"
)

var DB gorm.DB

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, _ := strconv.Atoi(id)

	var user = model.Account{ID: aid}
	DB.First(&user)
	ctx.JSON(http.StatusOK, user)
}

// ListAccounts godoc
//
//	@Summary		List accounts
//	@Description	get accounts
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"	Format(email)
//	@Success		200	{array}		model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	var accounts []model.Account
	DB.Find(&accounts)
	ctx.JSON(http.StatusOK, accounts)
}

// AddAccount godoc
//
//	@Summary		Add an account
//	@Description	add by json account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body		model.AddAccount	true	"Add account"
//	@Success		200		{object}	model.Account
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
	var Account model.Account
	account := ctx.ShouldBindJSON(&Account)

	DB.Create(&model.Account{Name: Account.Name})

	ctx.JSON(http.StatusOK, account)
}

// UpdateAccount godoc
//
//	@Summary		Update an account
//	@Description	Update by json account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Account ID"
//	@Param			account	body		model.UpdateAccount	true	"Update account"
//	@Success		200		{object}	model.Account
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/accounts/{id} [patch]
func (c *Controller) UpdateAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var Account model.Account
	account := ctx.ShouldBindJSON(&Account)
	// var updateAccount model.UpdateAccount
	// if err := ctx.ShouldBindJSON(&updateAccount); err != nil {
	// 	httputil.NewError(ctx, http.StatusBadRequest, err)
	// 	return
	// }
	// account := model.Account{
	// 	ID:   aid,
	// 	Name: updateAccount.Name,
	// }
	// err = account.Update()
	// if err != nil {
	// 	httputil.NewError(ctx, http.StatusNotFound, err)
	// 	return
	// }
	DB.Model(&Account).Where("id", aid).Update("Name", Account.Name)
	ctx.JSON(http.StatusOK, account)
}

// DeleteAccount godoc
//
//	@Summary		Delete an account
//	@Description	Delete by account ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"	Format(int64)
//	@Success		204	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [delete]
func (c *Controller) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	// err = model.Delete(aid)
	// if err != nil {
	// 	httputil.NewError(ctx, http.StatusNotFound, err)
	// 	return
	// }
	var Account model.Account
	DB.Unscoped().Delete(&Account, aid)
	ctx.JSON(http.StatusNoContent, gin.H{})
}
