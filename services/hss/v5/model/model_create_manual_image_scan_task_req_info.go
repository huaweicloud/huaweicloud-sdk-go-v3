package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManualImageScanTaskReqInfo 创建手动扫描任务请求体
type CreateManualImageScanTaskReqInfo struct {

	// **参数解释**: 扫描风险类型 **约束限制**: 不涉及 **取值范围**: - 0：none。 - 0x7fffffff：全部。 - 0x000f0000：漏洞。 - 0x0000f000：基线检查。 - 0x00000f00：恶意文件。 - 0x000000f0：敏感信息。 - 0x0000000f：软件合规。  **默认取值**: 不涉及
	ScanScope *int32 `json:"scan_scope,omitempty"`

	// **参数解释**: 扫描限速 单位：个/h **约束限制**: 不涉及 **取值范围**: 0-1000，0表示不限制。  **默认取值**: 不涉及
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// **参数解释**: 扫描全部镜像 **约束限制**: 不涉及 **取值范围**: - true：扫描全部镜像。 - false：指定镜像扫描,见image_info字段。  **默认取值**: 不涉及
	IsAll *bool `json:"is_all,omitempty"`

	QueryInfo *CreateManualImageScanTaskReqInfoQueryInfo `json:"query_info,omitempty"`

	// 待扫描镜像
	ImageInfo *[]CreateManualImageScanTaskReqInfoImageInfo `json:"image_info,omitempty"`
}

func (o CreateManualImageScanTaskReqInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualImageScanTaskReqInfo struct{}"
	}

	return strings.Join([]string{"CreateManualImageScanTaskReqInfo", string(data)}, " ")
}
