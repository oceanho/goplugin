package model

import "plugin"

// Plugin Author info
type PluginAuthor struct {
	Name  string
	Email string
}

type PluginExtra struct {
	plugin.Plugin
	Type   string
	Author PluginAuthor
}

type PluginList struct {
	Plugins []*PluginExtra
}

type PluginContext struct {

} 

type PluginCommon interface {
	GetAuthor() PluginAuthor
	GetPluginInfo() PluginExtra
	Call(context *PluginContext) (interface{}, error)
}
