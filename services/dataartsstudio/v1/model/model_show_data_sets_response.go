package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataSetsResponse Response Object
type ShowDataSetsResponse struct {

	// 查询结果总量
	Count float32 `json:"count,omitempty"`

	// 资产实体列表
	Entities *[]Entity `json:"entities,omitempty"`

	// 资产分类facets维度信息列表，数据结构List<Map<String, List<Aggregation>>> 取值为count
	Facets *interface{} `json:"facets,omitempty"`

	// 资产分类metrics维度信息列表，数据结构List<Map<String, List<Aggregation>>>  取值为aggregation
	Metrics *interface{} `json:"metrics,omitempty"`

	// 关联资产，数据类型Map<String, Entity>
	ReferredEntities *interface{} `json:"referred_entities,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ShowDataSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataSetsResponse struct{}"
	}

	return strings.Join([]string{"ShowDataSetsResponse", string(data)}, " ")
}
