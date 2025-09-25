package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageScanPolicyResponse Response Object
type ListImageScanPolicyResponse struct {

	// 扫描全部镜像 | true:扫描全部镜像 false:指定镜像扫描
	IsAll *bool `json:"is_all,omitempty"`

	// 定时扫描策略开关
	Enabled *bool `json:"enabled,omitempty"`

	// 定时扫描周期 | 3 每三天 7 每一周 14 每两周
	ScanCycle *int32 `json:"scan_cycle,omitempty"`

	// 扫描风险类型 | 0:无 0x0fffff:全部 0x0f0000:漏洞 0x00f000:基线检查 0x000f00:恶意文件 0x0000f0:敏感信息 0x00000f:软件合规
	ScanScope *int32 `json:"scan_scope,omitempty"`

	// 扫描限速 单位：个/h | 0 不限制
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// 时间范围 单位：天 | 0 全部镜像 1 3 7
	TimeScope *int32 `json:"time_scope,omitempty"`

	// 镜像仓库列表总数
	TotalRegistryNum *int32 `json:"total_registry_num,omitempty"`

	// 镜像列表
	RegistryList   *[]ImageScanPolicyRespInfoRegistryList `json:"registry_list,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ListImageScanPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageScanPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListImageScanPolicyResponse", string(data)}, " ")
}
