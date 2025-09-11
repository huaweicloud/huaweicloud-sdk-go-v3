package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorResource struct {
	FlavorId *string `json:"flavor_id,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o FlavorResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorResource struct{}"
	}

	return strings.Join([]string{"FlavorResource", string(data)}, " ")
}
