package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CveInfo struct {

	// **参数解释** CVE id **取值范围** 字符长度0-128位
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释** CVSS分数 **取值范围** 取值0-100
	CvssScore *float32 `json:"cvss_score,omitempty"`

	// **参数解释** CVE公布时间，时间单位：毫秒（ms） **取值范围** 取值0-3857960855552
	PublishTime *int64 `json:"publish_time,omitempty"`

	// **参数解释** CVE描述 **取值范围** 字符长度0-256位
	Description *string `json:"description,omitempty"`
}

func (o CveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CveInfo struct{}"
	}

	return strings.Join([]string{"CveInfo", string(data)}, " ")
}
