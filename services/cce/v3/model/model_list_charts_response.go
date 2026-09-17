package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListChartsResponse Response Object
type ListChartsResponse struct {

	// **参数解释：** 模板列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Body           *[]ChartResp `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListChartsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChartsResponse struct{}"
	}

	return strings.Join([]string{"ListChartsResponse", string(data)}, " ")
}
