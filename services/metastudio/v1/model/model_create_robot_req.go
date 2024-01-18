package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRobotReq 创建应用请求。
type CreateRobotReq struct {

	// 应用名称。
	Name string `json:"name"`

	// 对接第三方应用厂商类型。 > 0：科大讯飞AIUI；1：华为云CBS；2：科大讯飞星火交互认知大模型；5：第三方驱动
	AppType int32 `json:"app_type"`

	// 对话的并发数
	Concurrency int32 `json:"concurrency"`

	HuaweiEiCbs *HuaweiEiCbs `json:"huawei_ei_cbs,omitempty"`

	IflytekAiuiConfig *IflytekAiuiConfig `json:"iflytek_aiui_config,omitempty"`

	IflytekSpark *IflytekSpark `json:"iflytek_spark,omitempty"`
}

func (o CreateRobotReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRobotReq struct{}"
	}

	return strings.Join([]string{"CreateRobotReq", string(data)}, " ")
}
