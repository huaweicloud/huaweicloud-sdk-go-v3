package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteInstancesRequest struct {
	// 是否批量删除创建失败的缓存实例。取值如下： - true，表示删除租户所有创建失败的缓存实例，此时请求参数instances可为空； - false或者其他值，表示删除instances参数数组中指定的缓存实例。

	AllFailure *bool `json:"all_failure,omitempty"`

	Body *BatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstancesRequest", string(data)}, " ")
}
