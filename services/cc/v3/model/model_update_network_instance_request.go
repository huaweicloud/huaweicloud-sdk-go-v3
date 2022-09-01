package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateNetworkInstanceRequest struct {

	// 网络实例ID。
	Id string `json:"id" xml:"id"`

	Body *UpdateNetworkInstanceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateNetworkInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateNetworkInstanceRequest", string(data)}, " ")
}
