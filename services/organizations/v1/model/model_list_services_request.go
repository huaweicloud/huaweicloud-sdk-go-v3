package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServicesRequest Request Object
type ListServicesRequest struct {
}

func (o ListServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesRequest struct{}"
	}

	return strings.Join([]string{"ListServicesRequest", string(data)}, " ")
}
