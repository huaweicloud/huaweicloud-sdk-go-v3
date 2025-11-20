package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConfigRequest Request Object
type ListGlobalConfigRequest struct {
}

func (o ListGlobalConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConfigRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalConfigRequest", string(data)}, " ")
}
