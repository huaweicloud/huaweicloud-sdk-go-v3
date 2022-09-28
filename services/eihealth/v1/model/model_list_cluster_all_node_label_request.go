package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClusterAllNodeLabelRequest struct {
}

func (o ListClusterAllNodeLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterAllNodeLabelRequest struct{}"
	}

	return strings.Join([]string{"ListClusterAllNodeLabelRequest", string(data)}, " ")
}
