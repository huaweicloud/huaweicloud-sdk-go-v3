package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserDnsMappingResponse Response Object
type UpdateUserDnsMappingResponse struct {
	Error *CommonResponseErrorOfApiTest `json:"error,omitempty"`

	// 错误原因
	Reason *string `json:"reason,omitempty"`

	Result *DnsMapping `json:"result,omitempty"`

	// 状态值，如success、error
	Status *string `json:"status,omitempty"`

	// 错误码
	Code           *string `json:"code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserDnsMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserDnsMappingResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserDnsMappingResponse", string(data)}, " ")
}
