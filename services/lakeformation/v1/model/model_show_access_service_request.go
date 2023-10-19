package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessServiceRequest Request Object
type ShowAccessServiceRequest struct {
}

func (o ShowAccessServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessServiceRequest struct{}"
	}

	return strings.Join([]string{"ShowAccessServiceRequest", string(data)}, " ")
}
