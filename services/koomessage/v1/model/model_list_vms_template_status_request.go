package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVmsTemplateStatusRequest struct {

	// 智能信息基础版模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息基础版模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 智能信息基础版模板分类。 - public：公共模板 - private：个人模板
	TplType *string `json:"tpl_type,omitempty"`

	// 智能信息基础版模板是否携带参数，不传查全部模板。
	HasParam *bool `json:"has_param,omitempty"`

	// 智能信息基础版模板创建开始时间。 样例为：2019-10-12T07:20:50Z。  > begin_time和end_time必须全部为空或全部不为空，并且begin_time不能大于end_time。
	BeginTime *string `json:"begin_time,omitempty"`

	// 智能信息基础版模板创建结束时间。 样例为：2019-10-12T07:20:50Z。  > begin_time和end_time必须全部为空或全部不为空，并且begin_time不能大于end_time。
	EndTime *string `json:"end_time,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset int32 `json:"offset"`

	// 每页显示的条目数量。
	Limit int32 `json:"limit"`
}

func (o ListVmsTemplateStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsTemplateStatusRequest struct{}"
	}

	return strings.Join([]string{"ListVmsTemplateStatusRequest", string(data)}, " ")
}
