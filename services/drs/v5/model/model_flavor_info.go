package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorInfo 规格信息
type FlavorInfo struct {

	// 规格id
	Id *string `json:"id,omitempty"`

	// 云服务类型
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 规格类型码
	SpecTypeCode *string `json:"spec_type_code,omitempty"`

	// 规格码
	SpecCode *string `json:"spec_code,omitempty"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
