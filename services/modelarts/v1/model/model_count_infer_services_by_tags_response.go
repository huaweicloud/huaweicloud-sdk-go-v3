package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountInferServicesByTagsResponse Response Object
type CountInferServicesByTagsResponse struct {

	// **参数解释：** 资源实例总数量。 **取值范围：** 不涉及
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountInferServicesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInferServicesByTagsResponse struct{}"
	}

	return strings.Join([]string{"CountInferServicesByTagsResponse", string(data)}, " ")
}
