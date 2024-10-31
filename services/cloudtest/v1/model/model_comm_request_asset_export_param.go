package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestAssetExportParam struct {
	Params *AssetExportParam `json:"params"`
}

func (o CommRequestAssetExportParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestAssetExportParam struct{}"
	}

	return strings.Join([]string{"CommRequestAssetExportParam", string(data)}, " ")
}
