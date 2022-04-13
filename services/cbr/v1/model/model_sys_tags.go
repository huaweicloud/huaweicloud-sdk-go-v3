package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type SysTags struct {
	// 键。  系统标签的key，从白名单中取，不能随意定义。 目前仅支持 _sys_enterprise_project_id字段，对应 的value为企业项目ID。

	Key string `json:"key"`
	// 值列表。  目前仅会用到企业项目ID，其中默 认的企业项目ID为“0”。

	Values []string `json:"values"`
}

func (o SysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTags struct{}"
	}

	return strings.Join([]string{"SysTags", string(data)}, " ")
}
