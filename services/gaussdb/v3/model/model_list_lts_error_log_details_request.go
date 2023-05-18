package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLtsErrorLogDetailsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *LtsLogErrorQueryRequest `json:"body,omitempty"`
}

func (o ListLtsErrorLogDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsErrorLogDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsErrorLogDetailsRequest", string(data)}, " ")
}
