package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginWebCliResponse Response Object
type LoginWebCliResponse struct {

	// 客户端ID
	ClientId *string `json:"client_id,omitempty"`

	// DB数量
	Databases      *int32 `json:"databases,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o LoginWebCliResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginWebCliResponse struct{}"
	}

	return strings.Join([]string{"LoginWebCliResponse", string(data)}, " ")
}
