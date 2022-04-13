package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultAddResourceReq struct {
	// 资源列表

	Resources []ResourceCreate `json:"resources"`
}

func (o VaultAddResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultAddResourceReq struct{}"
	}

	return strings.Join([]string{"VaultAddResourceReq", string(data)}, " ")
}
