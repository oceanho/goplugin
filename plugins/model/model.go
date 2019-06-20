// Package model defined all of interface

// import "github.com/oceanho/goplugin/plugins/model"
package model

import "plugin"

// PluginAuthor defined a author info of Plugin
type PluginAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PluginExtra
type PluginExtra struct {
	plugin.Plugin
	Type   string       `json:"type"`
	ID     string       `json:"id"`
	Author PluginAuthor `json:"author"`
}

// PluginList
type PluginList struct {
	Plugins []*PluginExtra
}

// PluginContext define a Context
type PluginContext struct {
}

// PluginCommon
type PluginCommon interface {
	GetAuthor() PluginAuthor
	GetPluginInfo() PluginExtra
	Call(context *PluginContext) (interface{}, error)
}
