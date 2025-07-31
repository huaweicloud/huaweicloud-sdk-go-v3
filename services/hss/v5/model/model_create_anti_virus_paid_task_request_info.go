package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAntiVirusPaidTaskRequestInfo 创建付费扫描任务
type CreateAntiVirusPaidTaskRequestInfo struct {

	// 任务名称
	TaskName string `json:"task_name"`

	// 任务类型，包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
	ScanType string `json:"scan_type"`

	// 处置动作，包含如下:   - auto ：自动处置   - manual : 人工处置
	Action string `json:"action"`

	// 病毒查杀主机列表
	HostIds []string `json:"host_ids"`

	// 文件类型集合型，包含如下:   - 0 ：全部   - 1 : 可执行   - 2 : 压缩   - 3 : 脚本   - 4 : 文档   - 5 : 图片   - 6 : 音视频
	FileTypes *[]int32 `json:"file_types,omitempty"`

	// 扫描目录，多个用;分隔
	ScanDir *string `json:"scan_dir,omitempty"`

	// 排除目录，多个用;分隔
	IgnoreDir *string `json:"ignore_dir,omitempty"`

	// 任务ID 创建病毒扫描任务时,task_id是null.重新扫描时，task_id不是null,是当前任务的ID
	TaskId *string `json:"task_id,omitempty"`
}

func (o CreateAntiVirusPaidTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiVirusPaidTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateAntiVirusPaidTaskRequestInfo", string(data)}, " ")
}
