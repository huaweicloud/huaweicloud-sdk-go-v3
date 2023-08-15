package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeModeRequest Request Object
type ChangeModeRequest struct {

	// 指定修改的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *ChangeModeReq `json:"body,omitempty"`
}

func (o ChangeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeModeRequest", string(data)}, " ")
}
