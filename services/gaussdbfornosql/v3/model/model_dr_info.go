package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DrInfo **参数解释：** 容灾信息。 **约束限制：** 创建容灾实例时必传。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type DrInfo struct {

	// **参数解释：** 容灾源实例ID。可以调用查询实例列表和详情-QueryingInstancesandDetails接口获取。 **约束限制：** - 创建容灾实例时该参数必传，表示为该源实例创建容灾实例。 - 源实例为GeminiDB Cassandra实例。 - 源实例状态为正常。 - 传该参数时，Datastore的type参数的值必须为“cassandra”。 - 一个源实例只能有一个容灾实例。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`
}

func (o DrInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DrInfo struct{}"
	}

	return strings.Join([]string{"DrInfo", string(data)}, " ")
}
