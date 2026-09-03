package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdvancedConfig struct {

	// 分块开关 1:打开 0：关闭，默认：打开
	BlockEnable *string `json:"blockEnable,omitempty"`

	// 用例超时时间
	CaseTimeout *int64 `json:"caseTimeout,omitempty"`

	// httpClient报存cookie配置：1 保存 0 不保存
	EnableCookie *string `json:"enableCookie,omitempty"`

	// 关闭默认添加content-type和accept请求头配置：1 打开 0 关闭
	HeaderDefault *string `json:"headerDefault,omitempty"`

	// http请求超时时间
	HttpTimeout *int64 `json:"httpTimeout,omitempty"`

	// 八爪鱼镜像地址
	OctopusImage *string `json:"octopusImage,omitempty"`

	// 并行用例个数
	ParallelNumber *int32 `json:"parallelNumber,omitempty"`

	// 代理用户名
	ProxyAuthName *string `json:"proxyAuthName,omitempty"`

	// 代理密码
	ProxyAuthPassword *string `json:"proxyAuthPassword,omitempty"`

	// 代理服务配置
	ProxyHostName *string `json:"proxyHostName,omitempty"`

	// 代理服务配置
	ProxyPort *string `json:"proxyPort,omitempty"`

	// 串行配置
	SerialRun *string `json:"serialRun,omitempty"`

	// 任务停止时间
	TaskStopTime *sdktime.SdkTime `json:"taskStopTime,omitempty"`

	// 任务超时时间
	TaskTimeout *int64 `json:"taskTimeout,omitempty"`
}

func (o AdvancedConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedConfig struct{}"
	}

	return strings.Join([]string{"AdvancedConfig", string(data)}, " ")
}
