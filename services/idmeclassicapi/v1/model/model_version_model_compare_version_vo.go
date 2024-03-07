package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelCompareVersionVo struct {

	// 基础版本号。
	BasicVersion string `json:"basicVersion"`

	// 对比版本号。
	CorrelationVersion string `json:"correlationVersion"`

	// 主对象ID。
	Id string `json:"id"`
}

func (o VersionModelCompareVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelCompareVersionVo struct{}"
	}

	return strings.Join([]string{"VersionModelCompareVersionVo", string(data)}, " ")
}
