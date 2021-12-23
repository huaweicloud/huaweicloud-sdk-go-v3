package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 服务发现规则命名部分。
type NameRule struct {
	// nameType取值cmdLine时args格式为[\"start\",\"end\"],表示抽取命令行中start、end之间的字符。 nameType取值cmdLine时args格式为[\"aa\"],表示抽取环境变量名为aa对应的环境变量值。 nameType取值str时,args格式为[\"fix\"],表示服务名称最后拼接固定文字fix。 nameType取值cmdLineHash时,args格式为[\"0001\"],value格式为[\"ser\"],表示当启动命令是0001时,服务名称为ser。 服务命名部分,数组中有多个对象时表示将每个对象抽取到的字符串拼接作为服务的名称。

	AppNameRule []AppNameRule `json:"appNameRule"`
	// nameType取值cmdLine时args格式为[\"start\",\"end\"],表示抽取命令行中start、end之间的字符。 nameType取值cmdLine时args格式为 [\"aa\"],表示抽取环境变量名为aa对应的环境变量值。 nameType取值str时,args格式为[\"fix\"],表示服务名称最后拼接固定文字fix。 nameType取值cmdLineHash时,args格式为[\"0001\"],value格式为[\"ser\"],表示当启动命令是0001时,应用名称为ser。 应用命名部分。

	ApplicationNameRule []ApplicationNameRule `json:"applicationNameRule"`
}

func (o NameRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameRule struct{}"
	}

	return strings.Join([]string{"NameRule", string(data)}, " ")
}
