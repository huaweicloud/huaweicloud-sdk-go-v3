package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkGaussMySqlProxyRequestBody 数据库代理节点缩容信息
type ShrinkGaussMySqlProxyRequestBody struct {

	// 数据库代理节点缩容操作需要减少的节点数。  缩容的节点数的取值范围：1~30之间的整数。  限制条件：该实例的数据库代理节点的总数量小于等于32，大于等于2。
	NodeNum int32 `json:"node_num"`
}

func (o ShrinkGaussMySqlProxyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkGaussMySqlProxyRequestBody struct{}"
	}

	return strings.Join([]string{"ShrinkGaussMySqlProxyRequestBody", string(data)}, " ")
}
