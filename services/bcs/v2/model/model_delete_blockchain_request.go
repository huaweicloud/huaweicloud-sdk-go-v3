package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteBlockchainRequest struct {
	// bcs 服务id

	BlockchainId string `json:"blockchain_id"`
	// [是否删除存储，IEF模式下不用填写](tag:online,g42)[是否删除存储](tag:hcs)

	IsDeleteStorage *bool `json:"is_delete_storage,omitempty"`
	// [是否删除obs，IEF模式下不用填写](tag:online,g42)[是否删除obs](tag:hcs)

	IsDeleteObs *bool `json:"is_delete_obs,omitempty"`
	// 是否删除底层CCE资源，IEF模式下不用填写

	IsDeleteResource *bool `json:"is_delete_resource,omitempty"`
}

func (o DeleteBlockchainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlockchainRequest struct{}"
	}

	return strings.Join([]string{"DeleteBlockchainRequest", string(data)}, " ")
}
