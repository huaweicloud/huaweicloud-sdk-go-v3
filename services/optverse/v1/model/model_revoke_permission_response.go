package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokePermissionResponse Response Object
type RevokePermissionResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o RevokePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokePermissionResponse struct{}"
	}

	return strings.Join([]string{"RevokePermissionResponse", string(data)}, " ")
}
