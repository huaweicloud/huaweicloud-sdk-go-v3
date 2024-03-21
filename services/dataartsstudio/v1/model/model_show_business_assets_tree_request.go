package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBusinessAssetsTreeRequest Request Object
type ShowBusinessAssetsTreeRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)
	Workspace string `json:"workspace"`

	// 资产guid，未填充时查询LV1级别业务资产,获取方法请参见[数据开发作业ID](dataartsstudio_02_0351.xml)
	Guid *string `json:"guid,omitempty"`
}

func (o ShowBusinessAssetsTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBusinessAssetsTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowBusinessAssetsTreeRequest", string(data)}, " ")
}
