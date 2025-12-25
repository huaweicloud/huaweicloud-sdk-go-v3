package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudLogAliasResponse Response Object
type ListCloudLogAliasResponse struct {

	// alias名称
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCloudLogAliasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudLogAliasResponse struct{}"
	}

	return strings.Join([]string{"ListCloudLogAliasResponse", string(data)}, " ")
}
