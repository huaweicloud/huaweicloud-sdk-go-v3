package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHarvestJobStatusRequest Request Object
type UpdateHarvestJobStatusRequest struct {

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-Internal访问服务。
	AccessControlAllowInternal *string `json:"Access-Control-Allow-Internal,omitempty"`

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-External访问服务。
	AccessControlAllowExternal *string `json:"Access-Control-Allow-External,omitempty"`

	Body *UpdateHarvestJobStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateHarvestJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHarvestJobStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateHarvestJobStatusRequest", string(data)}, " ")
}
