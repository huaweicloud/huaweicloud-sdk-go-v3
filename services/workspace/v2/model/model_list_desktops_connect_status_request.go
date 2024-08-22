package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsConnectStatusRequest Request Object
type ListDesktopsConnectStatusRequest struct {

	// 桌面所属用户，批量筛选，最多不超过100个用户。
	UserNames *[]string `json:"user_names,omitempty"`

	// 桌面的登录状态。  - UNREGISTER：表示桌面未注册时的状态（桌面启动后，会自动注册）。关机后也会出现未注册的状态。 - REGISTERED：表示桌面注册以后，等待用户连接的状态。 - CONNECTED：表示用户已经成功登录，正在使用桌面。 - DISCONNECTED：表示桌面与客户端断开会话后显示的状态，可能为关闭客户端窗口，或客户端与桌面网络断开引起。
	ConnectStatus *string `json:"connect_status,omitempty"`

	// 从查询结果中的第几条数据开始返回,用于分页查询，取值范围0-2000，默认从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果中想要返回的信息条目数量,用于分页查询，取值范围0-2000，默认值100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopsConnectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsConnectStatusRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsConnectStatusRequest", string(data)}, " ")
}
