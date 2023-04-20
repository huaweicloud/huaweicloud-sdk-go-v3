package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAppInstanceRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`

	Body *CreateAppInstanceRequestDto `json:"body,omitempty"`
}

func (o CreateAppInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateAppInstanceRequest", string(data)}, " ")
}
