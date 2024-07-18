package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HmVulnInfoData 数据
type HmVulnInfoData struct {

	// 报告
	Report *[]HmVulnInfoDataReport `json:"report,omitempty"`
}

func (o HmVulnInfoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HmVulnInfoData struct{}"
	}

	return strings.Join([]string{"HmVulnInfoData", string(data)}, " ")
}
