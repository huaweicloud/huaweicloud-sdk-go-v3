package model

import (
	"encoding/json"

	"strings"
)

type VaultRemoveResourceReq struct {
	// 要移除的资源ID列表

	ResourceIds []string `json:"resource_ids"`
}

func (o VaultRemoveResourceReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultRemoveResourceReq struct{}"
	}

	return strings.Join([]string{"VaultRemoveResourceReq", string(data)}, " ")
}
