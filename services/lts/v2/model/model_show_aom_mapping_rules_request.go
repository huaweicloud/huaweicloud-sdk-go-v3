package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAomMappingRulesRequest Request Object
type ShowAomMappingRulesRequest struct {
}

func (o ShowAomMappingRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAomMappingRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowAomMappingRulesRequest", string(data)}, " ")
}
