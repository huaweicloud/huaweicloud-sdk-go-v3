package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobTagsResponse Response Object
type ListTrainingJobTagsResponse struct {

	// **参数解释**：标签列表，按key聚合，每个key下包含该项目下该key出现过的所有不同value。 **取值范围**：不涉及。
	Tags           *[]ProjectTag `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListTrainingJobTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTrainingJobTagsResponse", string(data)}, " ")
}
