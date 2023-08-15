package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsEipsRequest Request Object
type ListDesktopsEipsRequest struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面所属用户。
	UserName *string `json:"user_name,omitempty"`

	// EIP地址。
	Address *string `json:"address,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。如果不指定，则返回所有符合条件的记录。
	Limit *int32 `json:"limit,omitempty"`

	// EIP绑定状态。 - bind：表示已绑定的EIP。 - unbind：表示未绑定的EIP。
	State *string `json:"state,omitempty"`
}

func (o ListDesktopsEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsEipsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsEipsRequest", string(data)}, " ")
}
