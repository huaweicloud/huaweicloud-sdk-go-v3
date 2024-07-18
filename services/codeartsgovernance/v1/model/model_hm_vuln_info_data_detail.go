package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HmVulnInfoDataDetail 报告详情
type HmVulnInfoDataDetail struct {

	// 问题
	VulnInfo *[]HmVulnInfoDataDetailVulnInfo `json:"vuln_info,omitempty"`
}

func (o HmVulnInfoDataDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HmVulnInfoDataDetail struct{}"
	}

	return strings.Join([]string{"HmVulnInfoDataDetail", string(data)}, " ")
}
