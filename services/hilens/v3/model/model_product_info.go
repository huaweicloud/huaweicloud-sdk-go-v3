package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfo struct {
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	ResourceType *string `json:"resource_type,omitempty"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
