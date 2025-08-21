package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupDeployKeysRequest Request Object
type ListGroupDeployKeysRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGroupDeployKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupDeployKeysRequest struct{}"
	}

	return strings.Join([]string{"ListGroupDeployKeysRequest", string(data)}, " ")
}
