package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAomMappingRulesRequest struct {
	Body *AomMappingRequestInfo `json:"body,omitempty"`
}

func (o UpdateAomMappingRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAomMappingRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAomMappingRulesRequest", string(data)}, " ")
}
