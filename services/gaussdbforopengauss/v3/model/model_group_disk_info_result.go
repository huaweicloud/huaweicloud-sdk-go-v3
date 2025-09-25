package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupDiskInfoResult struct {

	// **参数解释**： 分片组ID。 **取值范围**： 不涉及。
	GroupId string `json:"group_id"`

	// **参数解释**： 分片名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 分片磁盘使用率。 **取值范围**： 不涉及。
	Used float64 `json:"used"`

	// **参数解释**： 分片磁盘大小。 **取值范围**： 不涉及。
	Total float64 `json:"total"`
}

func (o GroupDiskInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupDiskInfoResult struct{}"
	}

	return strings.Join([]string{"GroupDiskInfoResult", string(data)}, " ")
}
