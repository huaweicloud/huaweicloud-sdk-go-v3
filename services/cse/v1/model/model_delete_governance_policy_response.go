package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGovernancePolicyResponse Response Object
type DeleteGovernancePolicyResponse struct {

	// 结果信息
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGovernancePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGovernancePolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteGovernancePolicyResponse", string(data)}, " ")
}
