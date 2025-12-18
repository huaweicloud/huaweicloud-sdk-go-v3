package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModifyHistoryResponse Response Object
type ShowModifyHistoryResponse struct {

	// 实例参数的修改历史列表。
	Histories *[]ConfigurationHistoryRsp `json:"histories,omitempty"`

	// **参数解释：** 参数修改历史记录总条数。 **约束限制：** 默认返回参数历史修改记录总条数。若为参数名搜索，返回符合要求的记录总条数。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowModifyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModifyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowModifyHistoryResponse", string(data)}, " ")
}
