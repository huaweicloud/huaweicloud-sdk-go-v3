package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAomMappingRulesRequest Request Object
type CreateAomMappingRulesRequest struct {

	// 是否开启自动映射
	IsBatch bool `json:"isBatch"`

	Body *AomMappingRequestInfo `json:"body,omitempty"`
}

func (o CreateAomMappingRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAomMappingRulesRequest struct{}"
	}

	return strings.Join([]string{"CreateAomMappingRulesRequest", string(data)}, " ")
}
