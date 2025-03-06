package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClustersRequest Request Object
type ShowClustersRequest struct {
}

func (o ShowClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClustersRequest struct{}"
	}

	return strings.Join([]string{"ShowClustersRequest", string(data)}, " ")
}
