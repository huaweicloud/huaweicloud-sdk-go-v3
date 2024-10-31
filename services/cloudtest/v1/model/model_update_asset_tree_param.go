package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAssetTreeParam struct {

	// 资产树id
	Id *string `json:"id,omitempty"`

	// 资产树目录名称
	Name *string `json:"name,omitempty"`
}

func (o UpdateAssetTreeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetTreeParam struct{}"
	}

	return strings.Join([]string{"UpdateAssetTreeParam", string(data)}, " ")
}
