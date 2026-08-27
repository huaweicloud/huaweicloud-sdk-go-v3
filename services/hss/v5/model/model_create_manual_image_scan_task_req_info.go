package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManualImageScanTaskReqInfo 创建手动扫描任务请求体
type CreateManualImageScanTaskReqInfo struct {

	// **参数解释**: 扫描风险类型 **约束限制**: 不涉及 **取值范围**: - 0：none。 - 0x7fffffff：全部。 - 0x000f0000：漏洞。 - 0x0000f000：基线检查。 - 0x00000f00：恶意文件。 - 0x000000f0：敏感信息。 - 0x0000000f：软件合规。  **默认取值**: 不涉及
	ScanScope *int32 `json:"scan_scope,omitempty"`

	// **参数解释**: 三方镜像仓扫描限速,其他镜像仓不生效 单位：个/h **约束限制**: 不涉及 **取值范围**: 0-1000，0表示不限制。  **默认取值**: 不涉及
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// **参数解释**: 扫描全部镜像 **约束限制**: 不涉及 **取值范围**: - true：扫描全部镜像。支持全部镜像扫描或者指定镜像仓类型扫描，若为指定镜像仓类型扫描，需要填写query_info的image_type类型。 - false：指定镜像扫描,需要填写详细的镜像信息image_info字段。 **默认取值**: 不涉及
	IsAll *bool `json:"is_all,omitempty"`

	QueryInfo *CreateManualImageScanTaskReqInfoQueryInfo `json:"query_info,omitempty"`

	// **参数解释**:   待扫描镜像，is_all为false需要填写；   若为仓库镜像，需要填写id，image_digest，namespace，image_name，image_version，image_version，registry_id，registry_name，registry_type，若为企业镜像，需要填写instance_id   若为本地镜像，需要填写image_id，image_name，image_version，registry_id，registry_name，registry_type
	ImageInfo *[]CreateManualImageScanTaskReqInfoImageInfo `json:"image_info,omitempty"`
}

func (o CreateManualImageScanTaskReqInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualImageScanTaskReqInfo struct{}"
	}

	return strings.Join([]string{"CreateManualImageScanTaskReqInfo", string(data)}, " ")
}
