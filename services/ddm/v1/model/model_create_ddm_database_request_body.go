package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmDatabaseRequestBody databases参数说明
type CreateDdmDatabaseRequestBody struct {

	// 逻辑库名称，需要满足以下条件：  - 长度为2-48个字符。 - 必须以小写字母开头。 - 可以包含小写字母、数字、下划线，不能包含大写字母和其它特殊字符。 - 禁用关键字：  \"information_schema\"、\"mysql\"、\"performance_schema\"、\"sys\"。
	Name string `json:"name"`

	// 逻辑库的拆分模式。 - cluster表示逻辑库是拆分模式。 - single表示逻辑库是非拆分模式。
	ShardMode string `json:"shard_mode"`

	// 同一种工作模式下逻辑库分片的数量。 - shard_unit不为空， shard_unit与关联DN数量的乘积 - shard_unit为空，大于关联的DN数量，小于等于关联DN数量*64。
	ShardNumber int32 `json:"shard_number"`

	// 逻辑库关联的DN信息,最大个数为256。
	DnInstances []DatabaseDnInstances `json:"dn_instances"`
}

func (o CreateDdmDatabaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmDatabaseRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDdmDatabaseRequestBody", string(data)}, " ")
}
