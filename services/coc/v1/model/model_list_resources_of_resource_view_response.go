package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesOfResourceViewResponse Response Object
type ListResourcesOfResourceViewResponse struct {

	// **参数解释：** 视图聚合的资源数据。 **取值范围：** 不涉及。
	Data           *[]ListViewResourceResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListResourcesOfResourceViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesOfResourceViewResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesOfResourceViewResponse", string(data)}, " ")
}
