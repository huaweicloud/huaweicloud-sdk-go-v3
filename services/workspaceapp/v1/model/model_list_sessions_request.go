package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionsRequest Request Object
type ListSessionsRequest struct {

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 搜索开始时间，以会话开始时间为条件查询，格式2024-02-27T03:47:51.182Z。
	QueryBeginTime string `json:"query_begin_time"`

	// 搜索结束时间，以会话开始时间为条件查询，格式2024-02-27T03:47:51.182Z。
	QueryEndTime string `json:"query_end_time"`

	// AppServer组ID。
	AppServerGroupId *string `json:"app_server_group_id,omitempty"`

	// 服务器IP。
	VmIp *string `json:"vm_ip,omitempty"`

	// 应用服务器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 应用状态： * `Active` - 会话当前处于活动状态，有用户登录并且正在使用。 * `Disconnected` - 用户已经登录但会话处于断开连接状态。 * `AppcInit` - 会话正在初始化。 * `SignedOut` - 会话已注销。 * `InitFail` - 会话初始化失败。
	SessionState *string `json:"session_state,omitempty"`

	// 会话是否创建成功,默认不填则查询全部 * 'true' - 会话创建成功 * 'false' - 会话创建失败
	IsSuccess *string `json:"is_success,omitempty"`
}

func (o ListSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionsRequest struct{}"
	}

	return strings.Join([]string{"ListSessionsRequest", string(data)}, " ")
}
