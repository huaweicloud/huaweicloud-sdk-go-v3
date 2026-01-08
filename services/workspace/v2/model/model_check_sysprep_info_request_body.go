package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckSysprepInfoRequestBody struct {

	// 操作的桌面ID列表。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`
}

func (o CheckSysprepInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSysprepInfoRequestBody struct{}"
	}

	return strings.Join([]string{"CheckSysprepInfoRequestBody", string(data)}, " ")
}
