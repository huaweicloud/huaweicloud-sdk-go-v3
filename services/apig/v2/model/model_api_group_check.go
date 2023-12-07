package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiGroupCheck struct {

	// 待校验的API分组名称
	GroupName string `json:"group_name"`

	// 集成应用ID  暂不支持
	RomaAppId *string `json:"roma_app_id,omitempty"`
}

func (o ApiGroupCheck) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiGroupCheck struct{}"
	}

	return strings.Join([]string{"ApiGroupCheck", string(data)}, " ")
}
