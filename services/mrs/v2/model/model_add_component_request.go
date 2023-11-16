package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddComponentRequest Request Object
type AddComponentRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`

	Body *AddComponentsReq `json:"body,omitempty"`
}

func (o AddComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddComponentRequest struct{}"
	}

	return strings.Join([]string{"AddComponentRequest", string(data)}, " ")
}
