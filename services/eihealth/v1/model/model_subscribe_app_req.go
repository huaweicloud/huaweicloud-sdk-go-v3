package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeAppReq 订阅应用请求体
type SubscribeAppReq struct {

	// 资产id。长度1-128，只能包含字母、数字、下划线和中划线
	AssetId string `json:"asset_id"`

	// 资产版本。长度1-128，字母或数字开头，后面跟小写字母、数字、小数点、斜杠、下划线或中划线
	AssetVersion string `json:"asset_version"`

	// 目标应用名称 取值范围：长度为[1,56]，以大小写字母开头，允许出现中划线(-)、下划线(_)、小写字母和数字，且必须以大小写字母或数字结尾。
	DestinationAppName string `json:"destination_app_name"`

	// 目标应用版本。取值范围：长度[1,24]，以小写字母或数字或大写字母开头，允许出现中划线，必须以小写字母或数字或大写字母结尾。
	DestinationAppVersion string `json:"destination_app_version"`
}

func (o SubscribeAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeAppReq struct{}"
	}

	return strings.Join([]string{"SubscribeAppReq", string(data)}, " ")
}
