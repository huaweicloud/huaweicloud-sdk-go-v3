package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEntityInfoByGuidRequest Request Object
type ShowEntityInfoByGuidRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 资产的guid
	Guid string `json:"guid"`
}

func (o ShowEntityInfoByGuidRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntityInfoByGuidRequest struct{}"
	}

	return strings.Join([]string{"ShowEntityInfoByGuidRequest", string(data)}, " ")
}
