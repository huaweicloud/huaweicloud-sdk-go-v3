package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultItem struct {

	// CBR存储库ID
	Id *string `json:"id,omitempty"`

	// CBR存储库名称
	Name *string `json:"name,omitempty"`
}

func (o VaultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultItem struct{}"
	}

	return strings.Join([]string{"VaultItem", string(data)}, " ")
}
