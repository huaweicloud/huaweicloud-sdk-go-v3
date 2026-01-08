package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportPolicyGroupsResponse Response Object
type ExportPolicyGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportPolicyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportPolicyGroupsResponse struct{}"
	}

	return strings.Join([]string{"ExportPolicyGroupsResponse", string(data)}, " ")
}
