package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAntiVirusResultRequest Request Object
type ExportAntiVirusResultRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset int32 `json:"offset"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 服务器公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 处置状态，包含如下:   - unhandled ：未处理   - handled : 已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 威胁等级，包含如下:   - Low : 低危   - Medium : 中危   - High : 高危   - Critical : 致命
	SeverityList *[]string `json:"severity_list,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 病毒名称
	MalwareName *string `json:"malware_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 导出条数
	ExportSize *int32 `json:"export_size,omitempty"`

	// 文件hash,当前为sha256
	FileHash *string `json:"file_hash,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 是否使用手动隔离按钮
	ManualIsolate *bool `json:"manual_isolate,omitempty"`

	Body *ExportAntiVirusResultRequestBody `json:"body,omitempty"`
}

func (o ExportAntiVirusResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAntiVirusResultRequest struct{}"
	}

	return strings.Join([]string{"ExportAntiVirusResultRequest", string(data)}, " ")
}
