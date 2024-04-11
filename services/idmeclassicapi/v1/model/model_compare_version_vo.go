package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareVersionVo struct {

	// 基础版本号。
	BasicVersion string `json:"basicVersion"`

	// 对比版本号。
	CorrelationVersion string `json:"correlationVersion"`

	// 实例ID。
	Id string `json:"id"`
}

func (o CompareVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionVo struct{}"
	}

	return strings.Join([]string{"CompareVersionVo", string(data)}, " ")
}
