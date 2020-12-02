package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
	"strings"
)

type TreeController struct {
	beego.Controller
}

func (c *TreeController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	PubInit(c.Controller, ctx, controllerName, actionName, app)
}

/**
获取根目录结构
*/
func (c *TreeController) Get() {
	files, e := fileSystem.ListRoot()
	if e != nil {
		ServeJSON(c.Controller, e)
	}
	ServeJSON(c.Controller, files)
}

/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *TreeController) ListSubTree() {
	fileDir := c.GetString("fileDir")
	fileName := c.GetString("fileName")
	root, _ := c.GetBool("root")
	var destPath string
	if root && fileDir == "" {
		destPath = tools.PathSeparator
	} else {
		if fileDir == tools.PathSeparator {
			destPath = fileDir + fileName
		} else {
			destPath = fileDir
			if fileName != "" {
				if !strings.HasSuffix(fileDir, tools.PathSeparator) {
					destPath = fileDir + tools.PathSeparator + fileName
				} else {
					destPath = fileDir + fileName
				}

			}
		}
	}
	isDir, _ := fileSystem.IsDir(destPath)
	if isDir {
		authorization := c.Ctx.Request.Header["Authorization"]
		trimPrefix := ""
		if len(authorization) > 0 {
			myCustomClaims, _ := tools.GetTokenInfo(authorization[0])
			userInfo := service.GetUser(myCustomClaims.Name)
			trimPrefix = tools.PathSeparator + userInfo["userFullName"].(string)
		}
		files, _ := fileSystem.ListDir(destPath, trimPrefix)
		ServeJSON(c.Controller, files)
	} else {
		c.Data["json"] = &models.Result{}
		c.ServeJSON()
	}
}
