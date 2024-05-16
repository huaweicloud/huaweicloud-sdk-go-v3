package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SrCreateInstanceRspInstanceEngine 引擎信息。
type SrCreateInstanceRspInstanceEngine struct {

	// 引擎类型。
	Type *string `json:"type,omitempty"`

	// 引擎大版本号。
	Version *string `json:"version,omitempty"`
}

func (o SrCreateInstanceRspInstanceEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrCreateInstanceRspInstanceEngine struct{}"
	}

	return strings.Join([]string{"SrCreateInstanceRspInstanceEngine", string(data)}, " ")
}
