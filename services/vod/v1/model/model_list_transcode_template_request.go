package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTranscodeTemplateRequest struct {

	// 模板id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 是否默认
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default"`

	// 偏移量。默认为0。指定group_id时该参数无效。<br/>
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页记录数。默认为10，范围[1,100]。指定group_id时该参数无效。<br/>
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 按照模板名和描述模糊查询。指定group_id时该参数无效。<br/>
	QueryString *string `json:"query_string,omitempty" xml:"query_string"`
}

func (o ListTranscodeTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodeTemplateRequest", string(data)}, " ")
}
