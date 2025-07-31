package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageAssetStatisticsResponse Response Object
type ShowImageAssetStatisticsResponse struct {

	// **参数解释**: 本地镜像数量 **取值范围**: 最小值0，最大值65535
	LocalNum *int32 `json:"local_num,omitempty"`

	// **参数解释**: cicd镜像数量 **取值范围**: 最小值0，最大值65535
	CicdNum *int32 `json:"cicd_num,omitempty"`

	// **参数解释**: 仓库镜像数量 **取值范围**: 最小值0，最大值65535
	RegistryNum    *int32 `json:"registry_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageAssetStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageAssetStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowImageAssetStatisticsResponse", string(data)}, " ")
}
