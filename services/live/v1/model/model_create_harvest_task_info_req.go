package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHarvestTaskInfoReq 创建Live2Vod任务
type CreateHarvestTaskInfoReq struct {

	// 频道推流域名
	Domain string `json:"domain"`

	// 组名或应用名
	AppName string `json:"app_name"`

	// 频道ID。频道唯一标识，为必填项。频道ID不建议输入下划线“_”，否则会影响转码和截图任务
	Id string `json:"id"`

	// 开始时间。Unix时间戳：单位是秒
	StartTime int32 `json:"start_time"`

	// 结束时间。Unix时间戳：单位是秒
	EndTime int32 `json:"end_time"`

	// 事件名称。必选配置
	EventName string `json:"event_name"`

	// 任务描述，可选配置
	TaskDesc *string `json:"task_desc,omitempty"`

	PackageInfo *VodPackageInfo `json:"package_info,omitempty"`
}

func (o CreateHarvestTaskInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHarvestTaskInfoReq struct{}"
	}

	return strings.Join([]string{"CreateHarvestTaskInfoReq", string(data)}, " ")
}
