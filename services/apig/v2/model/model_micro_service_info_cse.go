package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CSE微服务详细信息
type MicroServiceInfoCse struct {

	// 微服务引擎编号
	EngineId string `json:"engine_id"`

	// 微服务编号
	ServiceId string `json:"service_id"`

	// 微服务引擎名称
	EngineName *string `json:"engine_name,omitempty"`

	// 微服务名称
	ServiceName *string `json:"service_name,omitempty"`

	// 注册中心地址
	RegisterAddress *string `json:"register_address,omitempty"`

	// 微服务所属的应用
	CseAppId *string `json:"cse_app_id,omitempty"`

	// 微服务的版本，已废弃，通过后端服务器组中的版本承载。
	Version *string `json:"version,omitempty"`
}

func (o MicroServiceInfoCse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCse struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCse", string(data)}, " ")
}
