package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionResponse Response Object
type ListPermissionResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionResponse", string(data)}, " ")
}
