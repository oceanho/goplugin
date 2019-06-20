package main

import (
	"fmt"
	"github.com/oceanho/goplugin/plugins/model"
	"github.com/oceanho/goplugin/utils/jsonlib"
	"plugin"
)

func main() {
	ldap, err := plugin.Open("build/plugins/auth/ldap.so")
	if err != nil {
		fmt.Errorf("load plugin faild. %v", err.Error())
	}
	sym, err := ldap.Lookup("PluginSymbol")
	if err != nil {
		fmt.Errorf("find plugin symbol fail. %v", err.Error())
	}
	ctx := &model.PluginContext{}
	pluginCommon := sym.(model.PluginCommon)
	pluginCommon.Call(ctx)
	fmt.Println(pluginCommon.GetAuthor())
	fmt.Println(pluginCommon.GetPluginInfo())

	author := pluginCommon.GetAuthor()
	jsonStr, err := jsonlib.Serializer(author)
	if err != nil {
		fmt.Errorf("serializer json fail, %v", err)
	}
	str := fmt.Sprintf("author json: %s", jsonStr)
	fmt.Println(str)
}
