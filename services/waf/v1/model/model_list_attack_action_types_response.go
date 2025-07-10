package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttackActionTypesResponse Response Object
type ListAttackActionTypesResponse struct {

	// CountItem的总数量
	Total *int32 `json:"total,omitempty"`

	// CountItem详细信息
	Items          *[]AttackActionCountItem `json:"items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAttackActionTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttackActionTypesResponse struct{}"
	}

	return strings.Join([]string{"ListAttackActionTypesResponse", string(data)}, " ")
}
