package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkipPreCheckInfo 跳过预检查项参数。
type SkipPreCheckInfo struct {

	// 跳过的预检查项。
	SkippedPrecheckList []string `json:"skipped_precheck_list"`

	// 跳过预检查原因。
	SkipReason string `json:"skip_reason"`
}

func (o SkipPreCheckInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkipPreCheckInfo struct{}"
	}

	return strings.Join([]string{"SkipPreCheckInfo", string(data)}, " ")
}
