package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHotInformationRequest Request Object
type ShowHotInformationRequest struct {
}

func (o ShowHotInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotInformationRequest struct{}"
	}

	return strings.Join([]string{"ShowHotInformationRequest", string(data)}, " ")
}
