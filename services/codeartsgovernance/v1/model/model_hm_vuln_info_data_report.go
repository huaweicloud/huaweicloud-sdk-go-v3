package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HmVulnInfoDataReport struct {
	Detail *HmVulnInfoDataDetail `json:"detail,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`
}

func (o HmVulnInfoDataReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HmVulnInfoDataReport struct{}"
	}

	return strings.Join([]string{"HmVulnInfoDataReport", string(data)}, " ")
}
