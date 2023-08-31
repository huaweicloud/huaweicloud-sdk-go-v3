package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIegPasswordResponse Response Object
type ChangeIegPasswordResponse struct {

	// 成功信息
	SuccessMsg     *string `json:"success_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeIegPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIegPasswordResponse struct{}"
	}

	return strings.Join([]string{"ChangeIegPasswordResponse", string(data)}, " ")
}
