package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceNamespaceResponse Response Object
type CreateInstanceNamespaceResponse struct {

	// 命名空间ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateInstanceNamespaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceNamespaceResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceNamespaceResponse", string(data)}, " ")
}
