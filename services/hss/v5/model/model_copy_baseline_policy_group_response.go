package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyBaselinePolicyGroupResponse Response Object
type CopyBaselinePolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CopyBaselinePolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyBaselinePolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"CopyBaselinePolicyGroupResponse", string(data)}, " ")
}
