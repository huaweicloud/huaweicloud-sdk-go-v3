package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAomMappingRulesRequest Request Object
type UpdateAomMappingRulesRequest struct {
	Body *UpdateAomMappingRequest `json:"body,omitempty"`
}

func (o UpdateAomMappingRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAomMappingRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAomMappingRulesRequest", string(data)}, " ")
}
