package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateUnbindPublicRequest struct {

	// 指定关闭公网访问集群ID。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *UnBindPublicReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateUnbindPublicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUnbindPublicRequest struct{}"
	}

	return strings.Join([]string{"UpdateUnbindPublicRequest", string(data)}, " ")
}
