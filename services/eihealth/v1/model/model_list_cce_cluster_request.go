package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceClusterRequest Request Object
type ListCceClusterRequest struct {
}

func (o ListCceClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceClusterRequest struct{}"
	}

	return strings.Join([]string{"ListCceClusterRequest", string(data)}, " ")
}
