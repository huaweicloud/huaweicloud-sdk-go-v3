package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferServicesByTagsResponse Response Object
type ListInferServicesByTagsResponse struct {

	// **参数解释：** 通过标签反查出来的资源列表。
	Resources *[]TmsResource `json:"resources,omitempty"`

	// **参数解释：** 总记录数。 **取值范围：** 不涉及
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInferServicesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferServicesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInferServicesByTagsResponse", string(data)}, " ")
}
