package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListNumberOfInstancesInDifferentStatusRequest struct {
	// 是否返回创建失败的实例数。   - 当参数值为“true”时，返回的统计包括创建失败的实例数。   - 参数值为“false”或者其他值，返回的统计不包括创建失败的实例数。

	IncludeFailure *string `json:"include_failure,omitempty"`
}

func (o ListNumberOfInstancesInDifferentStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNumberOfInstancesInDifferentStatusRequest struct{}"
	}

	return strings.Join([]string{"ListNumberOfInstancesInDifferentStatusRequest", string(data)}, " ")
}
