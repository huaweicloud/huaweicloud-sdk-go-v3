package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizePermissionResponse Response Object
type AuthorizePermissionResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o AuthorizePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizePermissionResponse struct{}"
	}

	return strings.Join([]string{"AuthorizePermissionResponse", string(data)}, " ")
}
