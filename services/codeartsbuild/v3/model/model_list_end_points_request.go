package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndPointsRequest Request Object
type ListEndPointsRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 每页显示的条目数量，limit小于等于100
	Limit int32 `json:"limit"`

	// **参数解释**： 分页页码，表示从此页开始查询。 **约束限制**： 不涉及。 **取值范围**： 只能使用数字，大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEndPointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndPointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndPointsRequest", string(data)}, " ")
}
