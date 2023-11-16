package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPolicyBlackAndWhiteIpListResponse Response Object
type AddPolicyBlackAndWhiteIpListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddPolicyBlackAndWhiteIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPolicyBlackAndWhiteIpListResponse struct{}"
	}

	return strings.Join([]string{"AddPolicyBlackAndWhiteIpListResponse", string(data)}, " ")
}
