package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserRequestDto 用户ID或第三方账号信息列表
type ShowUserRequestDto struct {

	// 用户ID或者第三方账号
	Id *string `json:"id,omitempty"`
}

func (o ShowUserRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserRequestDto struct{}"
	}

	return strings.Join([]string{"ShowUserRequestDto", string(data)}, " ")
}
