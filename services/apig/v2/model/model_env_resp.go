package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

type EnvResp struct {
	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 环境名称
	Name *string `json:"name,omitempty"`
	// 描述信息
	Remark *string `json:"remark,omitempty"`
	// 环境id
	Id *string `json:"id,omitempty"`
}

func (o EnvResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnvResp struct{}"
	}

	return strings.Join([]string{"EnvResp", string(data)}, " ")
}
