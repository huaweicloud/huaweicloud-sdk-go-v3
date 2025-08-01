package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateNameServer struct {

	// **参数解释：** 优先级。如果priority的值为“1”，表示会第一个采用该域名服务器进行解析。 **取值范围：** 不涉及。
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释：** DNS服务器地址。 **取值范围：** 不涉及。
	Address *string `json:"address,omitempty"`
}

func (o PrivateNameServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateNameServer struct{}"
	}

	return strings.Join([]string{"PrivateNameServer", string(data)}, " ")
}
