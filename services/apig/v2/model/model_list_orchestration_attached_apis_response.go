package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrchestrationAttachedApisResponse Response Object
type ListOrchestrationAttachedApisResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 编排规则绑定的API列表。
	Apis           *[]OrchestrationApiInfo `json:"apis,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListOrchestrationAttachedApisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrchestrationAttachedApisResponse struct{}"
	}

	return strings.Join([]string{"ListOrchestrationAttachedApisResponse", string(data)}, " ")
}
