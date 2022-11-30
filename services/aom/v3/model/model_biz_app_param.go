package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BizAppParam struct {

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 应用名称
	DisplayName *string `json:"display_name,omitempty"`

	// 应用关联的企业项目id
	EpsId *string `json:"eps_id,omitempty"`

	// 唯一标识
	Name string `json:"name"`

	// 前端默认是CONSOLE，不需要传参。rest接口无参数是API，有参数只能是：SERVICE_DISCOVERY
	RegisterType *string `json:"register_type,omitempty"`
}

func (o BizAppParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizAppParam struct{}"
	}

	return strings.Join([]string{"BizAppParam", string(data)}, " ")
}
