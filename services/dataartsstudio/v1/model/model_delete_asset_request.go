package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetRequest Request Object
type DeleteAssetRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 资产的guid
	Guid string `json:"guid"`
}

func (o DeleteAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetRequest", string(data)}, " ")
}
