package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询资产属性最新值请求
type LastAssetPropertyValueRequest struct {
	// 对property按指定tags标签进行过滤查询，填入资产标签属性的属性名与属性值，不可为空，例如 {\"tagPropertyA\": \"id0001\"}；注意，标签过滤只对打上标签时刻之后的数据生效，打标签之前的数据不能通过标签过滤

	Tags map[string]interface{} `json:"tags,omitempty"`
	// 属性过滤器，最多5个

	PropertyFilter *[]PropertyFilter `json:"property_filter,omitempty"`
	// 待查询的资产属性名列表,不携带该字段表示查询全部

	PropertyNames *[]string `json:"property_names,omitempty"`
}

func (o LastAssetPropertyValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LastAssetPropertyValueRequest struct{}"
	}

	return strings.Join([]string{"LastAssetPropertyValueRequest", string(data)}, " ")
}
