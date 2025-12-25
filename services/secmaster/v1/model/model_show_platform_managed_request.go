package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlatformManagedRequest Request Object
type ShowPlatformManagedRequest struct {
}

func (o ShowPlatformManagedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlatformManagedRequest struct{}"
	}

	return strings.Join([]string{"ShowPlatformManagedRequest", string(data)}, " ")
}
