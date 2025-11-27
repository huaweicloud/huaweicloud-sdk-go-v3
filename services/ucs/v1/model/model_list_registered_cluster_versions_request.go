package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegisteredClusterVersionsRequest Request Object
type ListRegisteredClusterVersionsRequest struct {
}

func (o ListRegisteredClusterVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegisteredClusterVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListRegisteredClusterVersionsRequest", string(data)}, " ")
}
