// Package main ldap provides LDAP authorization ability
package main

import (
	"fmt"
	"github.com/oceanho/goplugin/plugins/model"
)

// AuthImplLDAP
type AuthImplLDAP struct {

}

// GetAuthor
func (AuthImplLDAP) GetAuthor() (model.PluginAuthor) {
	author := new(model.PluginAuthor)
	author.Name = "ocean"
	author.Email = "gzhehai@foxmail.com"
	return *author
}

// GetPluginInfo
func (AuthImplLDAP) GetPluginInfo() (model.PluginExtra){
	pluginExtra := model.PluginExtra{}
	return pluginExtra
}

// Call
func (AuthImplLDAP) Call(context *model.PluginContext) (interface{}, error){
	fmt.Printf("ldap plugin")
	return nil,nil
}

// PluginSymbol
var PluginSymbol AuthImplLDAP