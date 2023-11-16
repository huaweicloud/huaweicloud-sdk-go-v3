package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyBlackAndWhiteIpListResponse Response Object
type DeletePolicyBlackAndWhiteIpListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyBlackAndWhiteIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyBlackAndWhiteIpListResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyBlackAndWhiteIpListResponse", string(data)}, " ")
}
