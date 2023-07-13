package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsSlowlogDetailsRequest Request Object
type ListLtsSlowlogDetailsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *LtsLogSlowQueryRequest `json:"body,omitempty"`
}

func (o ListLtsSlowlogDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowlogDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsSlowlogDetailsRequest", string(data)}, " ")
}
