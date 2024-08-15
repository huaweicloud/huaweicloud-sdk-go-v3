package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDetailRequest Request Object
type ShowInstanceDetailRequest struct {

	// 实例ID，获取方法请参见[获取集群列表](ListClusters.xml)。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceDetailRequest", string(data)}, " ")
}
