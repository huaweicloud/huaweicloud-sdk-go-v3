package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemediationRunRequestBody 手动执行修正的范围。
type RemediationRunRequestBody struct {

	// 是否选择对所有不合规资源执行补救。
	AllSupported bool `json:"all_supported"`

	ResourceIds *[]string `json:"resource_ids,omitempty"`
}

func (o RemediationRunRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationRunRequestBody struct{}"
	}

	return strings.Join([]string{"RemediationRunRequestBody", string(data)}, " ")
}
