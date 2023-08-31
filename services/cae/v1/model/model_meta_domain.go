package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetaDomain struct {

	// 域名ID。
	Id *string `json:"id,omitempty"`

	// 域名名称。
	Name *string `json:"name,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o MetaDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaDomain struct{}"
	}

	return strings.Join([]string{"MetaDomain", string(data)}, " ")
}
