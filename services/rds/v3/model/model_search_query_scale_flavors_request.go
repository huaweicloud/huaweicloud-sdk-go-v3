package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchQueryScaleFlavorsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o SearchQueryScaleFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryScaleFlavorsRequest struct{}"
	}

	return strings.Join([]string{"SearchQueryScaleFlavorsRequest", string(data)}, " ")
}
