package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetadataDeleteReq struct {

	// **参数解释**： 需要删除的任务列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及 **默认取值**： 不涉及。
	TaskIds []string `json:"task_ids"`
}

func (o MetadataDeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataDeleteReq struct{}"
	}

	return strings.Join([]string{"MetadataDeleteReq", string(data)}, " ")
}
