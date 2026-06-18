package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSwitchConfigInfoRequest Request Object
type ListSwitchConfigInfoRequest struct {
}

func (o ListSwitchConfigInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSwitchConfigInfoRequest struct{}"
	}

	return strings.Join([]string{"ListSwitchConfigInfoRequest", string(data)}, " ")
}
