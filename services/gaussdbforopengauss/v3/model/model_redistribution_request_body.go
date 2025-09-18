package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RedistributionRequestBody struct {

	// **参数解释**: 具有JOIN关系的表，指定该参数则启用多表扩容模式，扩容前设置生效。 如果指定过该参数，后续调用可以传入空数组清除多表扩容配置。 按照“database名称、schema1名称、table1名称、schema2名称、table2名称...”的格式指定，带有大小写或特殊字符的表名需要加\"\"转义。 多个数组则表示存在多个join group。 **约束限制**: 本次扩容结束后自动清除该配置，下次扩容需要重新设置。
	RedisJoinTables *[][]string `json:"redis_join_tables,omitempty"`

	// **参数解释**: 重分布并发数。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	RedisParallelJobs *int32 `json:"redis_parallel_jobs,omitempty"`

	// **参数解释**: 重分布资源管控级别。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	RedisResourceLevel *string `json:"redis_resource_level,omitempty"`
}

func (o RedistributionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedistributionRequestBody struct{}"
	}

	return strings.Join([]string{"RedistributionRequestBody", string(data)}, " ")
}
