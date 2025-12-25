package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchConditionResponse Response Object
type CreateSearchConditionResponse struct {

	// 检索条件ID
	ConditionId    *string `json:"condition_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSearchConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchConditionResponse struct{}"
	}

	return strings.Join([]string{"CreateSearchConditionResponse", string(data)}, " ")
}
