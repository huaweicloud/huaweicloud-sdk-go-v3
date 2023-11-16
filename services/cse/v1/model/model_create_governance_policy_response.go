package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGovernancePolicyResponse Response Object
type CreateGovernancePolicyResponse struct {

	// 结果信息
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGovernancePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGovernancePolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateGovernancePolicyResponse", string(data)}, " ")
}
