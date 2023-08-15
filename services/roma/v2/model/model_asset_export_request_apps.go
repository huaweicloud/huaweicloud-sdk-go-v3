package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetExportRequestApps struct {

	// 应用ID
	Id *string `json:"id,omitempty"`
}

func (o AssetExportRequestApps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetExportRequestApps struct{}"
	}

	return strings.Join([]string{"AssetExportRequestApps", string(data)}, " ")
}
