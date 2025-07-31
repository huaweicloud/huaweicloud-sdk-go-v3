package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddPolicyRequestInfo struct {

	// 检测特性规则列表
	FeatureList *[]ChkFeatureInfo `json:"feature_list,omitempty"`
}

func (o AddPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"AddPolicyRequestInfo", string(data)}, " ")
}
