package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MicroServiceInfoCseBase CSE微服务详细信息，service_type为CSE时必填
type MicroServiceInfoCseBase struct {

	// 微服务引擎编号
	EngineId string `json:"engine_id"`

	// 微服务编号
	ServiceId string `json:"service_id"`
}

func (o MicroServiceInfoCseBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoCseBase struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoCseBase", string(data)}, " ")
}
