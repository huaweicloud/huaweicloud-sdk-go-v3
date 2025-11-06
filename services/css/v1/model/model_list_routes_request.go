package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRoutesRequest Request Object
type ListRoutesRequest struct {

	// 指定待操作的集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o ListRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutesRequest struct{}"
	}

	return strings.Join([]string{"ListRoutesRequest", string(data)}, " ")
}
