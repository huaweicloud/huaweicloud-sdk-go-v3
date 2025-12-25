package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSearchConditionResponse Response Object
type UpdateSearchConditionResponse struct {

	// 检索条件ID
	ConditionId    *string `json:"condition_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSearchConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSearchConditionResponse struct{}"
	}

	return strings.Join([]string{"UpdateSearchConditionResponse", string(data)}, " ")
}
