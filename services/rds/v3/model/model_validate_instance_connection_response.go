package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateInstanceConnectionResponse Response Object
type ValidateInstanceConnectionResponse struct {

	// 响应结果
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateInstanceConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateInstanceConnectionResponse struct{}"
	}

	return strings.Join([]string{"ValidateInstanceConnectionResponse", string(data)}, " ")
}
