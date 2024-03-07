package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FindingId 访问分析结果的唯一标识符。
type FindingId struct {
}

func (o FindingId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FindingId struct{}"
	}

	return strings.Join([]string{"FindingId", string(data)}, " ")
}
