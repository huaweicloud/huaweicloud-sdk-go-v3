package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkGaussMySqlProxyRequestBody 数据库代理节点缩容信息
type ShrinkGaussMySqlProxyRequestBody struct {

	// 数据库代理节点缩容操作需要减少的节点数。  缩容的节点数的取值范围：1~30之间的整数。  限制条件：该实例的数据库代理节点的总数量小于等于32，大于等于2。
	NodeNum int32 `json:"node_num"`

	// **参数解释**：  数据库代理节点的节点ID。  获取方式请参见[查询数据库代理信息列表](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlProxyList.html)。  **约束限制**：  不传该字段，将随机删除代理节点；传入该字段，将删除指定ID的代理节点。
	NodeIds *[]string `json:"node_ids,omitempty"`
}

func (o ShrinkGaussMySqlProxyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkGaussMySqlProxyRequestBody struct{}"
	}

	return strings.Join([]string{"ShrinkGaussMySqlProxyRequestBody", string(data)}, " ")
}
