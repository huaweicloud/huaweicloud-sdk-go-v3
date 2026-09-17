package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceBody 更改绑定资源。
type UpdateResourceBody struct {

	// 需要更改绑定的集群id。
	ClusterId string `json:"cluster_id"`

	// 操作类型, 绑定:bind、解绑:unbind。
	Action string `json:"action"`
}

func (o UpdateResourceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceBody", string(data)}, " ")
}
