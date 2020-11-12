package controller

import (
	"errors"
	"github.com/weblair/lair/pkg/model"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weblair/lair/pkg/database"
)

type Controller interface {
	ResourceModel() model.Model
	LookupParam() string
	Create(c *gin.Context)
	List(c *gin.Context)
	Fetch(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func CreateRecord(ctx *gin.Context, ctl Controller) {
	logrus.WithFields(logrus.Fields{
		"resource_method": "create",
		"controller":      reflect.TypeOf(ctl).String(),
	}).Info("Attempting to create new record in database")

	ri := model.GetResourceModelInfo(ctl.ResourceModel())
	r := reflect.New(ri.ResourceType)

	err := ctx.Bind(r)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "request body is malformed",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

// ListRecords uses reflection to fetch a list of model instances from the database.
func ListRecords(ctx *gin.Context, ctl Controller, restrict bool) {
	// TODO: Restrict records by ownership
	// TODO: Panic if the model does not contain an ownership field and restrict is true
	// TODO: Respond with 204 for empty result sets

	logrus.WithFields(logrus.Fields{
		"resource_method": "list",
		"controller":      reflect.TypeOf(ctl).String(),
	}).Info("Fetching list of resources from database")

	s := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(ctl.ResourceModel())), 0, 0)
	p := reflect.New(s.Type())
	p.Elem().Set(s)
	r, ok := p.Interface().([]model.Model)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"resource_method": "list",
			"controller":      reflect.TypeOf(ctl).String(),
			"resource_type":      reflect.TypeOf(ctl.ResourceModel()).String(),
			"resource_list_type": reflect.TypeOf(r).String(),
		}).Error("Reflection did not return a slice of a model")

		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	logrus.WithFields(logrus.Fields{
		"resource_method": "list",
		"controller":      reflect.TypeOf(ctl).String(),
		"resource_type":      reflect.TypeOf(ctl.ResourceModel()).String(),
		"resource_list_type": reflect.TypeOf(r).String(),
	}).Debug("Resource reflection info")

	database.DB.List(&r)

	ctx.JSON(http.StatusOK, r)
}

func FetchRecord(c *gin.Context, ctl Controller, preload []string) {
	logrus.WithFields(logrus.Fields{
		"resource_method": "fetch",
		"controller":      reflect.TypeOf(ctl).String(),
	}).Info("Attempting to fetch from database")

	m := ctl.ResourceModel()
	param := ctl.LookupParam()

	err := database.DB.Fetch(&m, param, c.Param(param), preload)
	if err != nil {
		var e *database.NotFoundError
		if errors.As(err, e) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "resource not found",
				"id": c.Param(param),
			})
			return
		} else {
			logrus.WithFields(logrus.Fields{
				"resource_method": "fetch",
				"controller":      reflect.TypeOf(ctl).String(),
				"error": err,
			}).Error("Unrecognized error returned from data layer")
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
	}

	c.JSON(http.StatusOK, m)
}

func UpdateRecord(c *gin.Context, ctl Controller) {

}

func DeleteRecord(c *gin.Context, ctl Controller) {

}

