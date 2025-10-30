package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoIncrementUsageRequest Request Object
type ListAutoIncrementUsageRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 使用语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListAutoIncrementUsageRequestBody `json:"body,omitempty"`
}

func (o ListAutoIncrementUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoIncrementUsageRequest struct{}"
	}

	return strings.Join([]string{"ListAutoIncrementUsageRequest", string(data)}, " ")
}
