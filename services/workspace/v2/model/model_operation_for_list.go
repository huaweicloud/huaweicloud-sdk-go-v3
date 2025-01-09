package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationForList struct {

	// 主键UUID。
	Id *string `json:"id,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 录屏记录ID。
	ScreenRecordId *string `json:"screen_record_id,omitempty"`

	// 用户名。
	Username *string `json:"username,omitempty"`

	// 事件类型。 - APP：应用监控。 - FILE：文件监控。 - REG：注册表监控。 - HDP：协议行为监控。
	EventType *string `json:"event_type,omitempty"`

	// 事件级别。 - INFO：提示。 - ALARM：告警。 - ERROR：异常。
	EventLevel *string `json:"event_level,omitempty"`

	// 事件ID。 - APP_START：应用程序启动 - APP_STOP：应用程序结束 - APP_CRASH：应用程序异常退出 - APP_HANG：应用程序无响应 - APP_INSTALL：应用安装 - APP_UNINSTALL：应用卸裁 - FILE_CREATE：文件创建 - FILE_DELETE：文件删除 - FILE_RENAME：文件改名 - REG_CREATE：注册表创建 - REG_CHANGE：注册表修改 - REG_DELETE：注册表删除 - REG_RENAME：注册表改名 - REG_SETVALUE：注册表设置值 - HDP_USB：USB重定向事件 - HDP_CLIPBOARD：剪切板操作 - HDP_INPUTIDLE：空闲无操作 - HDP_PRINT：文件打印
	EventId *string `json:"event_id,omitempty"`

	// 事件内容。
	EventData *string `json:"event_data,omitempty"`

	// 操作时间（2024-10-15T11:04:41.263Z）。
	OperationTime *string `json:"operation_time,omitempty"`

	// 相对于视频开始的时间（秒）。
	RelativeStartTime *int32 `json:"relative_start_time,omitempty"`
}

func (o OperationForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationForList struct{}"
	}

	return strings.Join([]string{"OperationForList", string(data)}, " ")
}
