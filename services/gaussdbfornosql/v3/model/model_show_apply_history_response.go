package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplyHistoryResponse Response Object
type ShowApplyHistoryResponse struct {

	// **参数解释：** 总记录数。 **取值范围：** 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释：** 参数组模板应用历史列表。 **取值范围：** 不涉及。
	Histories      *[]ApplyHistoryRsp `json:"histories,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowApplyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowApplyHistoryResponse", string(data)}, " ")
}
