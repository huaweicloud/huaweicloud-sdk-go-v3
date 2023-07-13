package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecaptureDetectResponseResultDetail struct {

	// 标签值。| original：原始图 recapture：翻拍图
	Label *string `json:"label,omitempty"`

	// 置信度，取值范围（0~1）。
	Confidence *string `json:"confidence,omitempty"`
}

func (o RecaptureDetectResponseResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecaptureDetectResponseResultDetail struct{}"
	}

	return strings.Join([]string{"RecaptureDetectResponseResultDetail", string(data)}, " ")
}
