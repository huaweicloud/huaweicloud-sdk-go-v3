package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiPolicyGroupsResponse Response Object
type DeleteAiPolicyGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAiPolicyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiPolicyGroupsResponse struct{}"
	}

	return strings.Join([]string{"DeleteAiPolicyGroupsResponse", string(data)}, " ")
}
