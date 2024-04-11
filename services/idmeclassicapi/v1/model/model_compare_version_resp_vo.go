package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareVersionRespVo struct {

	// 基础版本对象。
	BasicVersion *interface{} `json:"basicVersion,omitempty"`

	// 当前版本对象。
	CorrelationVersion *interface{} `json:"correlationVersion,omitempty"`
}

func (o CompareVersionRespVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionRespVo struct{}"
	}

	return strings.Join([]string{"CompareVersionRespVo", string(data)}, " ")
}
