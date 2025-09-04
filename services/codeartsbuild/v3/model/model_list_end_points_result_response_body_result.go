package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndPointsResultResponseBodyResult 返回结果
type ListEndPointsResultResponseBodyResult struct {

	// **参数解释**： 私有仓库列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	Endpoints *[]EndPointResponse `json:"endpoints,omitempty"`

	// **参数解释**： 返回数据量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o ListEndPointsResultResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndPointsResultResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ListEndPointsResultResponseBodyResult", string(data)}, " ")
}
