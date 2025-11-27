package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClusterGroupDescriptionRequest struct {

	// 容器舰队描述信息
	Description string `json:"description"`
}

func (o UpdateClusterGroupDescriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupDescriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupDescriptionRequest", string(data)}, " ")
}
