package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareParameterGroupsRequest Request Object
type CompareParameterGroupsRequest struct {
	Body *ConfigurationDiffReqV3 `json:"body,omitempty"`
}

func (o CompareParameterGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareParameterGroupsRequest struct{}"
	}

	return strings.Join([]string{"CompareParameterGroupsRequest", string(data)}, " ")
}
