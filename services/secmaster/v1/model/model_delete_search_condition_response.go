package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSearchConditionResponse Response Object
type DeleteSearchConditionResponse struct {

	// 检索条件ID
	ConditionId    *string `json:"condition_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSearchConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSearchConditionResponse struct{}"
	}

	return strings.Join([]string{"DeleteSearchConditionResponse", string(data)}, " ")
}
