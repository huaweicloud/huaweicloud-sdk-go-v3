package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyBlackAndWhiteIpListRequest Request Object
type DeletePolicyBlackAndWhiteIpListRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *BlackWhiteIpRequestBody `json:"body,omitempty"`
}

func (o DeletePolicyBlackAndWhiteIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyBlackAndWhiteIpListRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyBlackAndWhiteIpListRequest", string(data)}, " ")
}
