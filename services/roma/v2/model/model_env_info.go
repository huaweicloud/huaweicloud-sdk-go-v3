package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvInfo struct {
	// 创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 环境名称

	Name *string `json:"name,omitempty"`
	// 描述信息

	Remark *string `json:"remark,omitempty"`
	// 环境编号

	Id *string `json:"id,omitempty"`
}

func (o EnvInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvInfo struct{}"
	}

	return strings.Join([]string{"EnvInfo", string(data)}, " ")
}
