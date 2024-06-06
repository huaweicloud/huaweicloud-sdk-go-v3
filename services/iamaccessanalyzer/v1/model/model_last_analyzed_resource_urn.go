package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LastAnalyzedResourceUrn 最近分析的资源的唯一资源标识符。
type LastAnalyzedResourceUrn struct {
}

func (o LastAnalyzedResourceUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LastAnalyzedResourceUrn struct{}"
	}

	return strings.Join([]string{"LastAnalyzedResourceUrn", string(data)}, " ")
}
