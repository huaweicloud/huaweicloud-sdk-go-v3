package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListIacFilesResponseInfoDataList struct {

	// **参数解释**: 文件ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64 **默认取值**: 不涉及
	FileId *string `json:"file_id,omitempty"`

	// **参数解释**: 文件名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-512 **默认取值**: 不涉及
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 文件类型 **约束限制**: 不涉及 **取值范围**: - dockerfile：Dockerfile文件 - k8s_yaml：k8s yaml文件  **默认取值**: 不涉及
	FileType *string `json:"file_type,omitempty"`

	// **参数解释**: 是否有风险 **约束限制**: 不涉及 **取值范围**: - true：有风险 - false：无风险 **默认取值**: 不涉及
	Risky *bool `json:"risky,omitempty"`

	// 风险项数量
	RiskNum *int32 `json:"risk_num,omitempty"`

	// 首次扫描时间
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// 最近扫描时间
	LastScanTime *int64 `json:"last_scan_time,omitempty"`

	// 上传文件时间
	UploadFileTime *int64 `json:"upload_file_time,omitempty"`

	// **参数解释**: cicd标识 **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	CicdId *string `json:"cicd_id,omitempty"`

	// **参数解释**: CI/CD名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	CicdName *string `json:"cicd_name,omitempty"`

	// **参数解释**: 扫描方式 **约束限制**: 不涉及 **取值范围**: - manual_scan：手动扫描 - cicd_scan：cicd扫描 **默认取值**: 不涉及
	ScanType *string `json:"scan_type,omitempty"`
}

func (o ListIacFilesResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFilesResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"ListIacFilesResponseInfoDataList", string(data)}, " ")
}
