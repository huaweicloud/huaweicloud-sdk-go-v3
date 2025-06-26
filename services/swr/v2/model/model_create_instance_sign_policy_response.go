package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceSignPolicyResponse Response Object
type CreateInstanceSignPolicyResponse struct {

	// 策略ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateInstanceSignPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceSignPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceSignPolicyResponse", string(data)}, " ")
}
