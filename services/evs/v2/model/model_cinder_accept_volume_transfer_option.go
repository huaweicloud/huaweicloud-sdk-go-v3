package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标记接受云硬盘过户操作。
type CinderAcceptVolumeTransferOption struct {

	// 云硬盘过户的身份认证密钥。  创建云硬盘过户时会返回该身份认证密钥。
	AuthKey string `json:"auth_key"`
}

func (o CinderAcceptVolumeTransferOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderAcceptVolumeTransferOption struct{}"
	}

	return strings.Join([]string{"CinderAcceptVolumeTransferOption", string(data)}, " ")
}
