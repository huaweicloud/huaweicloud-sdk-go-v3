package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionApplicationsRequest Request Object
type ListFunctionApplicationsRequest struct {
}

func (o ListFunctionApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionApplicationsRequest", string(data)}, " ")
}
