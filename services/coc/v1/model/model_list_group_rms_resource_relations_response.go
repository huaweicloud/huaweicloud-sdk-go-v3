package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupRmsResourceRelationsResponse Response Object
type ListGroupRmsResourceRelationsResponse struct {

	// **参数解释：** 分组资源关联信息。 **取值范围：** 长度大小在0~500之间。
	Data           *[]GroupRmsResourceRelationQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ListGroupRmsResourceRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupRmsResourceRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupRmsResourceRelationsResponse", string(data)}, " ")
}
