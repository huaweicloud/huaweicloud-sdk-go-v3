package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDesktopsDetailRequest struct {

	// 桌面状态。  - ACTIVE：运行中。 - SHUTOFF：关机。 - ERROR：异常。
	Status *string `json:"status,omitempty"`

	// 桌面所属用户。
	UserName *string `json:"user_name,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面IP地址。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-500，默认值500。
	Limit *int32 `json:"limit,omitempty"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面类型。  - DEDICATED：专属桌面。
	DesktopType *string `json:"desktop_type,omitempty"`

	// 桌面的标签。样例：  - key1=value1。 - key1=value1，key2=value2。
	Tag *string `json:"tag,omitempty"`
}

func (o ListDesktopsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsDetailRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsDetailRequest", string(data)}, " ")
}
