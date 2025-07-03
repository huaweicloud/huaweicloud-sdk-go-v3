package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfsRequest Request Object
type ListConfsRequest struct {

	// 指定查询集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定查询起始值，默认值为1。
	Start *string `json:"start,omitempty"`

	// 指定查询个数，默认值为10。
	Limit *string `json:"limit,omitempty"`
}

func (o ListConfsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfsRequest struct{}"
	}

	return strings.Join([]string{"ListConfsRequest", string(data)}, " ")
}
