package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDefectsResponse Response Object
type ShowTaskDefectsResponse struct {

	// **参数解释：** 缺陷详情列表，注意这里和入参中的limit数据不一定一致，服务会重新对响应数据拆分。
	Defects *[]DefectInfoV2 `json:"defects,omitempty"`

	// **参数解释：** 缺陷详情总数。 **取值范围：** 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskDefectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDefectsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDefectsResponse", string(data)}, " ")
}
