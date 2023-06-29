package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceLockRequest Request Object
type ShowWorkspaceLockRequest struct {
}

func (o ShowWorkspaceLockRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceLockRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceLockRequest", string(data)}, " ")
}
