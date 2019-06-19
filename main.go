package main

import (
	"fmt"
	"github.com/oceanho/goplugin/plugins/model"
	"plugin"
)

func main() {
	ldap, err := plugin.Open("build/plugins/auth/ldap.so")
	if err != nil {
		fmt.Errorf("load plugin faild. %v", err.Error())
	}
	symb, err := ldap.Lookup("PluginSymbol")
	if err != nil {
		fmt.Errorf("find plugin symbol fail. %v", err.Error())
	}
	ctx := &model.PluginContext{}
	pluginCommon := symb.(model.PluginCommon)
	pluginCommon.Call(ctx)
	fmt.Println(pluginCommon.GetAuthor())
	fmt.Println(pluginCommon.GetPluginInfo())
}
