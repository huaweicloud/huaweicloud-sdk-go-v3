package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePaginationResponseEntityV5 struct {

	// **参数解释**: 总记录数。 **取值范围**: 不涉及。
	TotalRecords *int32 `json:"total_records,omitempty"`

	// **参数解释**: 总页数。 **取值范围**: 不涉及。
	TotalPages *int32 `json:"total_pages,omitempty"`
}

func (o BasePaginationResponseEntityV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePaginationResponseEntityV5 struct{}"
	}

	return strings.Join([]string{"BasePaginationResponseEntityV5", string(data)}, " ")
}
