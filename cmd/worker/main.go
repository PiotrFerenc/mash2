package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()
	server.POST("/execute/:action", func(context *gin.Context) {
		name := context.Param("action")
		var parameters ActionContext

		if err := context.BindJSON(&parameters); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := Executor(name, parameters)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
	})

	err := server.Run("0.0.0.0:5000")
	if err != nil {
		return
	}
}

func Executor(actionName string, parameters ActionContext) error {
	actions := map[string]Action{
		"hallo": Hallo{name: "Word"},
	}
	action, exist := actions[actionName]
	if !exist {
		return fmt.Errorf("action %v not found", actionName)
	}
	action.Execute(parameters)
	return nil
}

type Action interface {
	Execute(parameters ActionContext) string
}

type Hallo struct {
	name string
}

func (receiver Hallo) Execute(parameters ActionContext) string {
	name := parameters.Parameters["name"]

	msg := "Hallo " + name
	log.Print(msg)
	return msg
}

type ActionContext struct {
	Parameters map[string]string `json:"parameters"`
}
