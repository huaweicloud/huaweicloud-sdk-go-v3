package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAspInfosResponse Response Object
type ListAspInfosResponse struct {

	// **参数解释**: asp报告信息列表。
	Asp *[]ListAspResult `json:"asp,omitempty"`

	// **参数解释**: 总记录数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAspInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAspInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAspInfosResponse", string(data)}, " ")
}
