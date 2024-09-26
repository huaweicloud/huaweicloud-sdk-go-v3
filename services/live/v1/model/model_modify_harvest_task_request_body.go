package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyHarvestTaskRequestBody 修改Live2Vod任务
type ModifyHarvestTaskRequestBody struct {

	// 频道推流域名
	Domain *string `json:"domain,omitempty"`

	// 组名或应用名
	AppName *string `json:"app_name,omitempty"`

	// 频道ID。频道唯一标识，为必填项。频道ID不建议输入下划线“_”，否则会影响转码和截图任务
	Id *string `json:"id,omitempty"`

	// 开始时间。Unix时间戳，单位为秒
	StartTime *int32 `json:"start_time,omitempty"`

	// 结束时间。Unix时间戳，单位为秒
	EndTime *int32 `json:"end_time,omitempty"`

	// 事件名称。必选配置
	EventName *string `json:"event_name,omitempty"`

	// 任务Id。必选配置
	JobId string `json:"job_id"`

	// 任务描述，可选配置
	TaskDesc *string `json:"task_desc,omitempty"`

	PackageInfo *VodPackageInfo `json:"package_info,omitempty"`
}

func (o ModifyHarvestTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHarvestTaskRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyHarvestTaskRequestBody", string(data)}, " ")
}
