package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityRiskResponseImageRiskRiskList struct {

	// **参数解释**： 风险级别 **取值范围**: 字符长度0-32位
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 镜像数量 **取值范围**: 最小值0，最大值2147483647
	ImageNum *int32 `json:"image_num,omitempty"`
}

func (o SecurityRiskResponseImageRiskRiskList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRiskResponseImageRiskRiskList struct{}"
	}

	return strings.Join([]string{"SecurityRiskResponseImageRiskRiskList", string(data)}, " ")
}
