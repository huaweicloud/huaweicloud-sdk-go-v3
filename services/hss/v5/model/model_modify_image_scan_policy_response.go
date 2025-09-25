package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyImageScanPolicyResponse Response Object
type ModifyImageScanPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyImageScanPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyImageScanPolicyResponse struct{}"
	}

	return strings.Join([]string{"ModifyImageScanPolicyResponse", string(data)}, " ")
}
