package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：platform_type为非DRIS时有效，表示第三方平台的对接参数。
type PlatformPara struct {

	// **参数说明**：第三方业务平台的ip地址和端口。
	Address *string `json:"address,omitempty"`

	// **参数说明**：鉴权用户名。
	Username *string `json:"username,omitempty"`

	// **参数说明**：鉴权密码，ITS800或者ATLAS500的密码
	Passwd *string `json:"passwd,omitempty"`
}

func (o PlatformPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlatformPara struct{}"
	}

	return strings.Join([]string{"PlatformPara", string(data)}, " ")
}
