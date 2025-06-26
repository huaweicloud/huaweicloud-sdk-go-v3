package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteInstanceSignPolicyResponse Response Object
type ExecuteInstanceSignPolicyResponse struct {

	// 执行ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExecuteInstanceSignPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteInstanceSignPolicyResponse struct{}"
	}

	return strings.Join([]string{"ExecuteInstanceSignPolicyResponse", string(data)}, " ")
}
