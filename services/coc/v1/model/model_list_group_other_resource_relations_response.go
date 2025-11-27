package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupOtherResourceRelationsResponse Response Object
type ListGroupOtherResourceRelationsResponse struct {

	// **参数解释：** 离线资源列表。 **取值范围：** 列表大小在0到500之间。
	Data           *[]GroupOtherResourceRelationQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o ListGroupOtherResourceRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupOtherResourceRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupOtherResourceRelationsResponse", string(data)}, " ")
}
