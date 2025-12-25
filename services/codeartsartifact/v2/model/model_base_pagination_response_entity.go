package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePaginationResponseEntity struct {

	// **参数解释**： 总记录数。 **取值范围**： 不涉及。
	TotalRecords *int32 `json:"totalRecords,omitempty"`

	// **参数解释**： 总页数。 **取值范围**： 不涉及。
	TotalPages *int32 `json:"totalPages,omitempty"`
}

func (o BasePaginationResponseEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePaginationResponseEntity struct{}"
	}

	return strings.Join([]string{"BasePaginationResponseEntity", string(data)}, " ")
}
