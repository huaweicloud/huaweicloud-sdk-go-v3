package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueryDetailRequest Request Object
type ShowQueryDetailRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	// 查询ID。
	QueryId string `json:"query_id"`

	// 采集时间。
	Ctime *int64 `json:"ctime,omitempty"`
}

func (o ShowQueryDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueryDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowQueryDetailRequest", string(data)}, " ")
}
