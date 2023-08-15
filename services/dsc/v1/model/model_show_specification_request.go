package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpecificationRequest Request Object
type ShowSpecificationRequest struct {
}

func (o ShowSpecificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecificationRequest struct{}"
	}

	return strings.Join([]string{"ShowSpecificationRequest", string(data)}, " ")
}
