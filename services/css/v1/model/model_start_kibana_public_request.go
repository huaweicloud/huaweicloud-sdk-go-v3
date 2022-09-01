package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartKibanaPublicRequest struct {

	// 指定待开启kibana公网访问的集群ID。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *StartKibanaPublicReq `json:"body,omitempty" xml:"body"`
}

func (o StartKibanaPublicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartKibanaPublicRequest struct{}"
	}

	return strings.Join([]string{"StartKibanaPublicRequest", string(data)}, " ")
}
