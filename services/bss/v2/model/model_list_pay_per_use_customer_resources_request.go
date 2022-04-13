package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPayPerUseCustomerResourcesRequest struct {
	Body *QueryResourcesReq `json:"body,omitempty"`
}

func (o ListPayPerUseCustomerResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPayPerUseCustomerResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListPayPerUseCustomerResourcesRequest", string(data)}, " ")
}
