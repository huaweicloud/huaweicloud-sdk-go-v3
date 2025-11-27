package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpecialUserRequest Request Object
type ShowSpecialUserRequest struct {
}

func (o ShowSpecialUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecialUserRequest struct{}"
	}

	return strings.Join([]string{"ShowSpecialUserRequest", string(data)}, " ")
}
