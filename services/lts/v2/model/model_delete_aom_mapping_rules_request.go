package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAomMappingRulesRequest struct {
	// 接入lts规则id

	Id string `json:"id"`
}

func (o DeleteAomMappingRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAomMappingRulesRequest struct{}"
	}

	return strings.Join([]string{"DeleteAomMappingRulesRequest", string(data)}, " ")
}
