package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetRepoRoleResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *Empty `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetRepoRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRepoRoleResponse struct{}"
	}

	return strings.Join([]string{"SetRepoRoleResponse", string(data)}, " ")
}
