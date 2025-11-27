package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExportTaskResponse Response Object
type CreateExportTaskResponse struct {

	// 导出任务唯一id。
	ExportTaskId *string `json:"export_task_id,omitempty"`

	// 导出源资源类型，支持批量任务导出类型BatchTask。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源过滤条件，Json格式，里面是(K,V)键值对，当resource_type为BatchTask时填写填写key为task_id，value为BatchTask的task_id(task_id从批量任务接口获得)。
	ResourceCondition *interface{} `json:"resource_condition,omitempty"`

	// 导出格式，目前xls格式。
	ExportFormat *string `json:"export_format,omitempty"`

	// 任务状态，状态分别有：Processing：执行中，Success：成功，Failed：失败。
	Status *string `json:"status,omitempty"`

	// 已导出的资源数量。
	ExportCount *int64 `json:"export_count,omitempty"`

	// 导出任务的执行进度，取值范围为1-100，当100表示进度为100%，任务完成。
	Progress *int32 `json:"progress,omitempty"`

	// 在物联网平台创建产品的时间，格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z'，如2020-08-12T12:12:12.333Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 导出任务的执行结束时间，格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z'，如2020-08-12T12:12:12.333Z。
	EndTime *string `json:"end_time,omitempty"`

	// 租户规则的生效范围，默认GLOBAL，取值如下： - GLOBAL：生效范围为租户级 - APP：生效范围为应用级，如果类型为APP，需要携带app_id，如果不带，生效范围为defaultApp。
	AppType *string `json:"app_type,omitempty"`

	// 应用ID。此参数为非必选参数，用于兼容平台老用户存在多应用的场景。存在多应用的用户需要使用该接口时，必须携带该参数指定创建的规则归属到哪个应用下，否则接口会提示错误。如果用户存在多应用，同时又不想携带该参数，可以联系华为技术支持对用户数据做应用合并。
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExportTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateExportTaskResponse", string(data)}, " ")
}
