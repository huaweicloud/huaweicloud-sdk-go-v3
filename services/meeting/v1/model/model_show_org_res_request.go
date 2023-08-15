package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrgResRequest Request Object
type ShowOrgResRequest struct {
}

func (o ShowOrgResRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrgResRequest struct{}"
	}

	return strings.Join([]string{"ShowOrgResRequest", string(data)}, " ")
}
