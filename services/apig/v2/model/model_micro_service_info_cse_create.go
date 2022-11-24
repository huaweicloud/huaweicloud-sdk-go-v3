package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MicroServiceInfoCseCreate struct {

	// 微服务引擎编号
	EngineId string `json:"engine_id"`

	// 微服务编号
	ServiceId string `json:"service_id"`

	// 微服务版本
	Version string `json:"version"`
}

func (o MicroServiceInfoCseCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCseCreate struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCseCreate", string(data)}, " ")
}
