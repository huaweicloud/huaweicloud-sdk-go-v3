package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersRequest Request Object
type ListClustersRequest struct {
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
