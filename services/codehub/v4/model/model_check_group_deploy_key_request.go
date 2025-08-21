package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckGroupDeployKeyRequest Request Object
type CheckGroupDeployKeyRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *DeployKeyValueDto `json:"body,omitempty"`
}

func (o CheckGroupDeployKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckGroupDeployKeyRequest struct{}"
	}

	return strings.Join([]string{"CheckGroupDeployKeyRequest", string(data)}, " ")
}
