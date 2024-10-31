package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestAssetTree struct {
	Params *AssetTree `json:"params,omitempty"`
}

func (o CommRequestAssetTree) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestAssetTree struct{}"
	}

	return strings.Join([]string{"CommRequestAssetTree", string(data)}, " ")
}
