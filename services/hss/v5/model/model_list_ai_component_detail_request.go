package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiComponentDetailRequest Request Object
type ListAiComponentDetailRequest struct {

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 资产类别 **约束限制**: 不涉及 **取值范围**: - host：主机资产 - container：容器资产  **默认取值**: host
	Category string `json:"category"`

	// **参数解释**: AI组件类别 **约束限制**: 不涉及 **取值范围**: - app：应用 - tool：工具  **默认取值**: 不涉及
	Catalogue string `json:"catalogue"`

	// **参数解释**: category==host时 表示服务器名称 category==container时 表示节点名称 category==serverless时 表示实例名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ServerName *string `json:"server_name,omitempty"`

	// **参数解释**: category==host时 表示服务器IP地址 category==container时 表示节点IP地址 category==serverless时 表示实例IP地址 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ServerIp *string `json:"server_ip,omitempty"`

	// **参数解释**: AI应用名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	AiApplication *string `json:"ai_application,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: AI工具名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	AiTool *string `json:"ai_tool,omitempty"`

	// **参数解释**: AI应用类型 **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**: AI版本 **约束限制**: 不涉及 **取值范围**: 字符长度0-32位 **默认取值**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**: 安装路径 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位 **默认取值**: 不涉及
	InstallationPath *string `json:"installation_path,omitempty"`

	// **参数解释**: 首次扫描时间，时间单位毫秒（ms） **约束限制**: 不涉及 **取值范围**: 取值0-9223372036854775807 **默认取值**: 不涉及
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// **参数解释**: 最近扫描时间，时间单位毫秒（ms） **约束限制**: 不涉及 **取值范围**: 取值0-9223372036854775807 **默认取值**: 不涉及
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 容器名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 容器ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`
}

func (o ListAiComponentDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiComponentDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAiComponentDetailRequest", string(data)}, " ")
}
