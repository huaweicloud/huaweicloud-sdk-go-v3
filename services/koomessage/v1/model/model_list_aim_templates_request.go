package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimTemplatesRequest Request Object
type ListAimTemplatesRequest struct {

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 模板分类。  - public：查公共模板 - private：查个人模板  > 不传查全部模板。
	TplType *string `json:"tpl_type,omitempty"`

	// 厂商类型。 - HUAWEI：华为 - XIAOMI：小米 - OPPO：OPPO - MEIZU：魅族 - VIVO：VIVO
	FactoryType *[]string `json:"factory_type,omitempty"`

	// 模板是否携带参数。  - true：是 - false：否
	HasParam *bool `json:"has_param,omitempty"`

	// 模板创建开始时间。样例：2019-10-12T07:20:50Z。  > begin_time和end_time必须全部为空或全部不为空，并且begin_time不能大于end_time。
	BeginTime *string `json:"begin_time,omitempty"`

	// 模板创建结束时间。样例：2019-10-12T07:20:50Z。  > begin_time和end_time必须全部为空或全部不为空，并且begin_time不能大于end_time。
	EndTime *string `json:"end_time,omitempty"`

	// 响应里只返回状态信息，不返回pages和params。  - false：默认值，返回全量信息 - true：只返回状态信息
	IsOnlyStatus *bool `json:"is_only_status,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAimTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAimTemplatesRequest", string(data)}, " ")
}
