package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CveInfo struct {

	// **参数解释**: CVE id **取值范围**: 字符长度0-64位
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: CVSS分数 **取值范围**: 最小值0，最大值10
	CvssScore *float32 `json:"cvss_score,omitempty"`

	// **参数解释**: CVE公布时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	PublishTime *int64 `json:"publish_time,omitempty"`

	// **参数解释**: CVE描述 **取值范围**: 字符长度0-4096位
	Description *string `json:"description,omitempty"`
}

func (o CveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CveInfo struct{}"
	}

	return strings.Join([]string{"CveInfo", string(data)}, " ")
}
