package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessageEmailConfigRequest Request Object
type ShowMessageEmailConfigRequest struct {
}

func (o ShowMessageEmailConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageEmailConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowMessageEmailConfigRequest", string(data)}, " ")
}
