package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupGeneralPolicyRequest Request Object
type UpdateGroupGeneralPolicyRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *UpdateGeneralPolicyDto `json:"body,omitempty"`
}

func (o UpdateGroupGeneralPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupGeneralPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupGeneralPolicyRequest", string(data)}, " ")
}
