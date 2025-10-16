package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteIpRequestBody struct {

	// **参数解释：** 需要删除IP的组类型。 **约束限制：** 不涉及。 **取值范围：** shard：表示删除所有Shard组IP。 config：表示删除Config组IP。 **默认取值：** 不涉及。
	Type string `json:"type"`
}

func (o DeleteIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteIpRequestBody", string(data)}, " ")
}
