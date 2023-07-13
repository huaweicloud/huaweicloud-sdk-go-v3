package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQualityTemplatesResponse Response Object
type ListQualityTemplatesResponse struct {

	// 总条数
	Count *int64 `json:"count,omitempty"`

	// 分页数据
	Resources      *[]RuleTemplateDetailVo `json:"resources,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListQualityTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListQualityTemplatesResponse", string(data)}, " ")
}
