package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetNameRequest struct {

	// 资产名称
	Name *string `json:"name,omitempty"`
}

func (o AssetNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetNameRequest struct{}"
	}

	return strings.Join([]string{"AssetNameRequest", string(data)}, " ")
}
