package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tsub/echo-sandbox/app/model"
)

func PostsIndex(c echo.Context) error {
	db, err := model.NewDatabase()
	if err != nil {
		return err
	}

	posts := model.AllPost(db.Conn)

	return c.JSON(http.StatusOK, posts)
}

func PostShow(c echo.Context) error {
	db, err := model.NewDatabase()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	post, err := model.FindPost(db.Conn, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, post)
}

func PostCreate(c echo.Context) error {
	db, err := model.NewDatabase()
	if err != nil {
		return err
	}

	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	post.Create(db.Conn)

	return c.JSON(http.StatusCreated, post)
}

func PostUpdate(c echo.Context) error {
	db, err := model.NewDatabase()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	post, err := model.FindPost(db.Conn, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	newPost := new(model.Post)
	if err := c.Bind(newPost); err != nil {
		return err
	}

	post.Update(db.Conn, newPost)

	return c.JSON(http.StatusOK, newPost)
}

func PostDelete(c echo.Context) error {
	db, err := model.NewDatabase()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	post, err := model.FindPost(db.Conn, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	post.Delete(db.Conn)

	return c.JSON(http.StatusOK, nil)
}
