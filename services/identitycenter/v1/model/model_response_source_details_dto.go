package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResponseSourceDetailsDto 联邦配置属性映射详细信息
type ResponseSourceDetailsDto struct {

	// 属性映射值
	Source []string `json:"source"`
}

func (o ResponseSourceDetailsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseSourceDetailsDto struct{}"
	}

	return strings.Join([]string{"ResponseSourceDetailsDto", string(data)}, " ")
}
