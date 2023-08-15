package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataCenterV2Do struct {

	// 区域类型0大陆 1国际
	Type *int32 `json:"type,omitempty"`

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 区域名称
	RegionName *string `json:"region_name,omitempty"`

	// 是否敏感
	IsSensitive *int32 `json:"is_sensitive,omitempty"`
}

func (o DataCenterV2Do) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataCenterV2Do struct{}"
	}

	return strings.Join([]string{"DataCenterV2Do", string(data)}, " ")
}
