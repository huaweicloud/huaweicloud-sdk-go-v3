package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesSeamlessOptions 通用音视频设置项。
type PoliciesSeamlessOptions struct {

	// 软件路径。长度不能超过1000个字符。
	SeamlessApplyPath *string `json:"seamless_apply_path,omitempty"`
}

func (o PoliciesSeamlessOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesSeamlessOptions struct{}"
	}

	return strings.Join([]string{"PoliciesSeamlessOptions", string(data)}, " ")
}
