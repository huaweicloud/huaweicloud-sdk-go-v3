package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupAliResourceRelationsResponse Response Object
type ListGroupAliResourceRelationsResponse struct {

	// **参数解释：** 阿里云资源列表。 **取值范围：** 列表大小在0到500之间，包含0和500。
	Data           *[]GroupAliResourceRelationQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ListGroupAliResourceRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupAliResourceRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupAliResourceRelationsResponse", string(data)}, " ")
}
