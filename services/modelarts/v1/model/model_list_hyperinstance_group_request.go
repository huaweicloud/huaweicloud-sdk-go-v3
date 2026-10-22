package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperinstanceGroupRequest Request Object
type ListHyperinstanceGroupRequest struct {
}

func (o ListHyperinstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListHyperinstanceGroupRequest", string(data)}, " ")
}
