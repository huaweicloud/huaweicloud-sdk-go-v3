package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBiactiveInstanceCandidatesResponse Response Object
type ListBiactiveInstanceCandidatesResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释：** 候选实例列表。 **取值范围：** 不涉及。
	Instances      *[]BiactiveInstanceV3 `json:"instances,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListBiactiveInstanceCandidatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBiactiveInstanceCandidatesResponse struct{}"
	}

	return strings.Join([]string{"ListBiactiveInstanceCandidatesResponse", string(data)}, " ")
}
