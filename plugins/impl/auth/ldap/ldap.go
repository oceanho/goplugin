package main

import (
	"fmt"
	"github.com/oceanho/goplugin/plugins/model"
)

type AuthImplLDAP struct {

}

func (AuthImplLDAP) GetAuthor() (model.PluginAuthor) {
	author := new(model.PluginAuthor)
	author.Name = "ocean"
	author.Email = "gzhehai@foxmail.com"
	return *author
}

func (AuthImplLDAP) GetPluginInfo() (model.PluginExtra){
	pluginExtra := model.PluginExtra{}
	return pluginExtra
}
func (AuthImplLDAP) Call(context *model.PluginContext) (interface{}, error){
	fmt.Printf("ldap plugin")
	return nil,nil
}

var PluginSymbol AuthImplLDAP