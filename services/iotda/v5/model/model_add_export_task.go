package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExportTask 创建导出任务结构体。
type AddExportTask struct {

	// 导出源资源类型，支持批量任务导出类型BatchTask。
	ResourceType string `json:"resource_type"`

	// 资源过滤条件，Json 格式，里面是(K,V)键值对，当resource_type为BatchTask时填写填写key为task_id，value为BatchTask的task_id(task_id从批量任务接口获得)。
	ResourceCondition *interface{} `json:"resource_condition"`

	// 导出格式，目前仅支持xls格式,默认格式为xls.(注意下载的文件已使用zip打包，后缀为'.zip'，此处格式指的导出内容承载格式)
	ExportFormat *string `json:"export_format,omitempty"`

	// 租户规则的生效范围，默认GLOBAL，取值如下： - GLOBAL：生效范围为租户级 - APP：生效范围为应用级，如果类型为APP，需要携带app_id，如果不带，生效范围为defaultApp。
	AppType *string `json:"app_type,omitempty"`

	// 应用ID。此参数为非必选参数，用于兼容平台老用户存在多应用的场景。存在多应用的用户需要使用该接口时，必须携带该参数指定创建的规则归属到哪个应用下，否则接口会提示错误。如果用户存在多应用，同时又不想携带该参数，可以联系华为技术支持对用户数据做应用合并。
	AppId *string `json:"app_id,omitempty"`
}

func (o AddExportTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExportTask struct{}"
	}

	return strings.Join([]string{"AddExportTask", string(data)}, " ")
}
