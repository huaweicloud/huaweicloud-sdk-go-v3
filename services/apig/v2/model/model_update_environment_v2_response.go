package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEnvironmentV2Response struct {

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 环境名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述信息
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 环境id
	Id             *string `json:"id,omitempty" xml:"id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnvironmentV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentV2Response struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentV2Response", string(data)}, " ")
}
