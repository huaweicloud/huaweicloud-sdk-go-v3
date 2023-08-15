package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplatesDetailResponse Response Object
type ShowTemplatesDetailResponse struct {

	// id
	Id *int32 `json:"id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// 目录ID
	CategoryId *int64 `json:"category_id,omitempty"`

	// Completeness:完整性,Uniqueness:唯一性,Timeliness:及时性,Validity:有效性,Accuracy:准确性,Consistency:一致性
	Dimension *string `json:"dimension,omitempty"`

	// 规则类型，Field:字段级规则,Table:表级规则,Database:库级规则,Cross-field:跨字段级规则,Customize:自定义规则
	Type *string `json:"type,omitempty"`

	// 是否为系统模板
	SystemTemplate *bool `json:"system_template,omitempty"`

	// 定义关系
	SqlInfo *string `json:"sql_info,omitempty"`

	// 结果说明
	ResultDescription *string `json:"result_description,omitempty"`

	// 创建时间,13位时间戳(精确到毫秒)
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者,System代表系统自带
	Creator        *string `json:"creator,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTemplatesDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplatesDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplatesDetailResponse", string(data)}, " ")
}
