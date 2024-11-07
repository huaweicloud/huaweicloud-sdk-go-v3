package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateRobotReq 校验应用请求。
type ValidateRobotReq struct {

	// 对接第三方应用厂商类型。 > 0：科大讯飞AIUI；1：华为云CBS；2：科大讯飞星火交互认知大模型；6：第三方语言模型；8：奇妙问
	AppType *int32 `json:"app_type,omitempty"`

	HuaweiEiCbs *HuaweiEiCbs `json:"huawei_ei_cbs,omitempty"`

	IflytekAiuiConfig *IflytekAiuiConfig `json:"iflytek_aiui_config,omitempty"`

	IflytekSpark *IflytekSpark `json:"iflytek_spark,omitempty"`

	ThirdPartyModelConfig *ThirdPartyModelConfig `json:"third_party_model_config,omitempty"`

	MobvoiConfig *MobvoiConfig `json:"mobvoi_config,omitempty"`
}

func (o ValidateRobotReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRobotReq struct{}"
	}

	return strings.Join([]string{"ValidateRobotReq", string(data)}, " ")
}
