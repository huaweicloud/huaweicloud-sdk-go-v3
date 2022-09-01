package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultRemoveResourceReq struct {

	// 要移除的资源ID列表
	ResourceIds []string `json:"resource_ids" xml:"resource_ids"`
}

func (o VaultRemoveResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultRemoveResourceReq struct{}"
	}

	return strings.Join([]string{"VaultRemoveResourceReq", string(data)}, " ")
}
