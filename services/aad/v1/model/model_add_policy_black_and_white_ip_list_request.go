package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPolicyBlackAndWhiteIpListRequest Request Object
type AddPolicyBlackAndWhiteIpListRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *BlackWhiteIpRequestBody `json:"body,omitempty"`
}

func (o AddPolicyBlackAndWhiteIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPolicyBlackAndWhiteIpListRequest struct{}"
	}

	return strings.Join([]string{"AddPolicyBlackAndWhiteIpListRequest", string(data)}, " ")
}
