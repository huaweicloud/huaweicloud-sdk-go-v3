package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessClientRequest Request Object
type UpdateAccessClientRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// 客户端ID。创建客户端时自动生成。
	ClientId string `json:"client_id"`

	Body *UpdateAccessClientRequestBody `json:"body,omitempty"`
}

func (o UpdateAccessClientRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessClientRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccessClientRequest", string(data)}, " ")
}
