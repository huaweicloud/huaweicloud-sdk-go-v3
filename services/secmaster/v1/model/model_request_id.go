package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequestId **参数解释**: 请求的ID **约束限制**: 不涉及
type RequestId struct {
}

func (o RequestId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestId struct{}"
	}

	return strings.Join([]string{"RequestId", string(data)}, " ")
}
