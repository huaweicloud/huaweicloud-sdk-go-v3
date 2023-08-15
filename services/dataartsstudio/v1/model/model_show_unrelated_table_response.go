package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUnrelatedTableResponse Response Object
type ShowUnrelatedTableResponse struct {
	Attributes *AttributeSearchResult `json:"attributes,omitempty"`

	// 分类
	Classification *string `json:"classification,omitempty"`

	// 结果总量
	Count *int32 `json:"count,omitempty"`

	// 资产信息
	Entities *[]AtlasEntityHeader `json:"entities,omitempty"`

	FullTextResult *[]AtlasFullTextResult `json:"full_text_result,omitempty"`

	// 查询内容
	QueryText *string `json:"query_text,omitempty"`

	// 查询类型，取值范围DSL,FULL_TEXT,GREMLIN,BASIC,ATTRIBUTE,RELATIONSHIP,ADVANCED
	QueryType *string `json:"query_type,omitempty"`

	// Map<String, AtlasEntityHeader>
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`

	// 滚动条id
	ScrollId *string `json:"scroll_id,omitempty"`

	// 参数
	SearchParameters *interface{} `json:"search_parameters,omitempty"`

	// 类型
	Type           *string `json:"type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUnrelatedTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUnrelatedTableResponse struct{}"
	}

	return strings.Join([]string{"ShowUnrelatedTableResponse", string(data)}, " ")
}
