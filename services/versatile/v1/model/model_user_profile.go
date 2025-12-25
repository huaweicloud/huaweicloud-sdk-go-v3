package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserProfile struct {

	// 运行时, 是否读取用户画像
	EnableRetrieve *bool `json:"enable_retrieve,omitempty"`

	// 运行时, 是否构建用户画像
	EnableExtract *bool `json:"enable_extract,omitempty"`
}

func (o UserProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserProfile struct{}"
	}

	return strings.Join([]string{"UserProfile", string(data)}, " ")
}
