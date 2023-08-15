package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuiltInConformancePackTemplatesResponse Response Object
type ListBuiltInConformancePackTemplatesResponse struct {

	// 预定义合规规则包模板列表。
	Value *[]ConformancePackTemplate `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBuiltInConformancePackTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuiltInConformancePackTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListBuiltInConformancePackTemplatesResponse", string(data)}, " ")
}
