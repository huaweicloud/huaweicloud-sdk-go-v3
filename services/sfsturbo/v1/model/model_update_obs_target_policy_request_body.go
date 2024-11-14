package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsTargetPolicyRequestBody 更新后端存储自动同步策略请求体
type UpdateObsTargetPolicyRequestBody struct {
	Policy *ObsDataRepositoryPolicy `json:"policy"`
}

func (o UpdateObsTargetPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsTargetPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateObsTargetPolicyRequestBody", string(data)}, " ")
}
