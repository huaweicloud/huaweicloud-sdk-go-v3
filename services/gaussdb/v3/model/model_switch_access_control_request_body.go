package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchAccessControlRequestBody struct {

	// 是否开启访问控制。 取值： - true：开启。 - false：关闭。
	OpenAccessControl bool `json:"open_access_control"`
}

func (o SwitchAccessControlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAccessControlRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchAccessControlRequestBody", string(data)}, " ")
}
