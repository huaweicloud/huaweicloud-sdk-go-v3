package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferIntranetConnectionApplicationsResponse Response Object
type ListInferIntranetConnectionApplicationsResponse struct {

	// **参数解释：** 当前页。 **取值范围：** 不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释：** 总页数。 **取值范围：** 不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释：** 每页大小。 **取值范围：** 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释：** 数据总量。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 申请信息列表。
	Data           *[]IntranetConnectionInfo `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListInferIntranetConnectionApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferIntranetConnectionApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListInferIntranetConnectionApplicationsResponse", string(data)}, " ")
}
