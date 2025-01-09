package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatMappingConfigsRequest Request Object
type UpdateNatMappingConfigsRequest struct {
	Body *UpdateNatMappingConfigsReq `json:"body,omitempty"`
}

func (o UpdateNatMappingConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatMappingConfigsRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatMappingConfigsRequest", string(data)}, " ")
}
