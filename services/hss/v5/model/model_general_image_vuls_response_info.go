package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneralImageVulsResponseInfo 所有镜像漏洞详情
type GeneralImageVulsResponseInfo struct {

	// 漏洞名称
	VulName *string `json:"vul_name,omitempty"`

	// 漏洞ID
	VulId *string `json:"vul_id,omitempty"`

	// 漏洞类型，包含如下：   -linux_vul : linux漏洞   -app_vul : 应用漏洞
	Type *string `json:"type,omitempty"`

	// 镜像类型，包含如下：   -local : 本地镜像   -registry : 仓库镜像   -cicd : CI/CD镜像   -cluster : 集群镜像
	ImageType *string `json:"image_type,omitempty"`

	// 漏洞标签列表
	LabelList *[]string `json:"label_list,omitempty"`

	// 漏洞的风险程度，取值如下：  -Critical : 严重  -High : 高危  -Medium : 中危  -Low : 低危
	SeverityLevel *string `json:"severity_level,omitempty"`

	// 受影响镜像总数
	ImageNum *int32 `json:"image_num,omitempty"`

	// CVE列表
	CveList *[]GeneralImageVulsResponseInfoCveList `json:"cve_list,omitempty"`

	// 镜像最大CVSS分值
	MaxCvssScore *float32 `json:"max_cvss_score,omitempty"`

	// 最近扫描时间，时间单位：毫秒（ms）
	ScanTime *int64 `json:"scan_time,omitempty"`

	// 漏洞描述
	Description *string `json:"description,omitempty"`

	// 漏洞修复参考链接
	Url *string `json:"url,omitempty"`

	// 修复建议
	SolutionDetail *string `json:"solution_detail,omitempty"`

	// 受影响集群总数
	ClusterNum *int32 `json:"cluster_num,omitempty"`
}

func (o GeneralImageVulsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralImageVulsResponseInfo struct{}"
	}

	return strings.Join([]string{"GeneralImageVulsResponseInfo", string(data)}, " ")
}
