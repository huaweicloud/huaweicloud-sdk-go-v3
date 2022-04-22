package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFlavorByTypeRequest struct {

	// 指定待更改的集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定待更改的集群节点类型。
	Types string `json:"types"`

	Body *UpdateFlavorReq `json:"body,omitempty"`
}

func (o UpdateFlavorByTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlavorByTypeRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlavorByTypeRequest", string(data)}, " ")
}
