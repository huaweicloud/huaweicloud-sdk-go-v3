package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiGroupCheck struct {
	// 待校验的API分组名称

	GroupName string `json:"group_name"`
	// 集成应用ID  校验分组在集成应用下是否重名时必填，不填写默认校验全局分组是否重名

	RomaAppId *string `json:"roma_app_id,omitempty"`
}

func (o ApiGroupCheck) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiGroupCheck struct{}"
	}

	return strings.Join([]string{"ApiGroupCheck", string(data)}, " ")
}
