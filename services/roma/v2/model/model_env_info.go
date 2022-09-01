package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvInfo struct {

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 环境名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述信息
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 环境编号
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o EnvInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvInfo struct{}"
	}

	return strings.Join([]string{"EnvInfo", string(data)}, " ")
}
