package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsStatusRequest Request Object
type ListDesktopsStatusRequest struct {
}

func (o ListDesktopsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsStatusRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsStatusRequest", string(data)}, " ")
}
