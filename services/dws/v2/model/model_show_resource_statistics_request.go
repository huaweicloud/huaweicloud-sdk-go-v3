package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceStatisticsRequest Request Object
type ShowResourceStatisticsRequest struct {

	// **参数解释**： 命名空间。 **约束限制**： 不涉及。 **取值范围**： 固定值dws。 **默认取值**： dws。
	Namespace *string `json:"namespace,omitempty"`
}

func (o ShowResourceStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceStatisticsRequest", string(data)}, " ")
}
