package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SysTagsRes struct {

	// 企业项目的key填：_sys_enterprise_project_id。
	Key *string `json:"key,omitempty"`

	// 企业项目的id。可以从企业项目获取。
	Value *string `json:"value,omitempty"`
}

func (o SysTagsRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTagsRes struct{}"
	}

	return strings.Join([]string{"SysTagsRes", string(data)}, " ")
}
