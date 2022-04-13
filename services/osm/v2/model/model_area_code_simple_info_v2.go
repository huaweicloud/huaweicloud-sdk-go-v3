package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AreaCodeSimpleInfoV2 struct {
	// 唯一id

	Id *int32 `json:"id,omitempty"`
	// 国家码

	AreaCode *string `json:"area_code,omitempty"`
	// 国家名称

	AreaName *string `json:"area_name,omitempty"`
}

func (o AreaCodeSimpleInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AreaCodeSimpleInfoV2 struct{}"
	}

	return strings.Join([]string{"AreaCodeSimpleInfoV2", string(data)}, " ")
}
