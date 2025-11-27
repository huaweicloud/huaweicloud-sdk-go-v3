package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyDefinitionsRequest Request Object
type ListPolicyDefinitionsRequest struct {
}

func (o ListPolicyDefinitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyDefinitionsRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyDefinitionsRequest", string(data)}, " ")
}
