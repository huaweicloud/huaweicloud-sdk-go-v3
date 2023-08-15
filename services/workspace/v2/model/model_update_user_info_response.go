package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserInfoResponse Response Object
type UpdateUserInfoResponse struct {

	// 用户id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserInfoResponse", string(data)}, " ")
}
