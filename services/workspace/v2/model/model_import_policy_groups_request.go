package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportPolicyGroupsRequest Request Object
type ImportPolicyGroupsRequest struct {
	Body *ImportPolicyGroupsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportPolicyGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportPolicyGroupsRequest struct{}"
	}

	return strings.Join([]string{"ImportPolicyGroupsRequest", string(data)}, " ")
}
