package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleBinPolicyRequestBody struct {
	Policy *RecycleBinPolicy `json:"policy,omitempty"`
}

func (o RecycleBinPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"RecycleBinPolicyRequestBody", string(data)}, " ")
}
