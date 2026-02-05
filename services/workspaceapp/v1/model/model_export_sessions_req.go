package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSessionsReq 导出用户会话列表请求体。
type ExportSessionsReq struct {

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 搜索开始时间，以会话开始时间为条件查询，只支持导出30天内数据，格式2024-02-27T03:47:51.182Z，参数中query_begin_time与query_end_time必须同时存在或都不存在，都不存在时导出近一个月的数据。
	QueryBeginTime *sdktime.SdkTime `json:"query_begin_time,omitempty"`

	// 搜索结束时间，以会话开始时间为条件查询，只支持导出30天内数据，格式2024-02-27T03:47:51.182Z，参数中query_begin_time与query_end_time必须同时存在或都不存在，都不存在时导出近一个月的数据。
	QueryEndTime *sdktime.SdkTime `json:"query_end_time,omitempty"`

	// AppServer组ID。
	AppServerGroupId *string `json:"app_server_group_id,omitempty"`

	// 服务器IP。
	VmIp *string `json:"vm_ip,omitempty"`

	// 客户端出口IP。
	PublicIp *string `json:"public_ip,omitempty"`

	// 应用服务器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 应用状态，支持查询多个，中间用英文逗号分隔： * `Active` - 会话当前处于活动状态，有用户登录并且正在使用。 * `Disconnected` - 用户已经登录但会话处于断开连接状态。 * `AppcInit` - 会话正在初始化。 * `SignedOut` - 会话已注销。 * `InitFail` - 会话初始化失败。
	SessionState *string `json:"session_state,omitempty"`

	// 会话是否创建成功,默认不填则查询全部 * 'true' - 会话创建成功 * 'false' - 会话创建失败
	IsSuccess *string `json:"is_success,omitempty"`
}

func (o ExportSessionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSessionsReq struct{}"
	}

	return strings.Join([]string{"ExportSessionsReq", string(data)}, " ")
}
