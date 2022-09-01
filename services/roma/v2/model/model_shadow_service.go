package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShadowService struct {

	// 服务ID
	ServiceId *int32 `json:"service_id,omitempty" xml:"service_id"`

	// 服务名称
	ServiceName *string `json:"service_name,omitempty" xml:"service_name"`

	// 影子值
	Properties *[]ShadowValue `json:"properties,omitempty" xml:"properties"`
}

func (o ShadowService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShadowService struct{}"
	}

	return strings.Join([]string{"ShadowService", string(data)}, " ")
}
