package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProduct struct {

	// 变更后的资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 变更后的商品规格编码
	ResourceSpecCode string `json:"resource_spec_code"`

	// 变更后的资源配额 如果operate_type为addition时，resource_size必须要大于原来的resource_id，decrease时要小于原来的resource_size，并且大于等于当前project下的ecs数量
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 要进行变更的资源ID
	ResourceId string `json:"resource_id"`
}

func (o UpdateProduct) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProduct struct{}"
	}

	return strings.Join([]string{"UpdateProduct", string(data)}, " ")
}
