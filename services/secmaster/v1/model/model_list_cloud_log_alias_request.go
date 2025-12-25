package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudLogAliasRequest Request Object
type ListCloudLogAliasRequest struct {
}

func (o ListCloudLogAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudLogAliasRequest struct{}"
	}

	return strings.Join([]string{"ListCloudLogAliasRequest", string(data)}, " ")
}
