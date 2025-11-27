package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityProfile struct {

	// 安全态势感知项配置值名称
	Key string `json:"key"`

	// 安全态势感知项配置值，数据格式参考创建安全态势感知接口说明
	Value *interface{} `json:"value"`
}

func (o SecurityProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityProfile struct{}"
	}

	return strings.Join([]string{"SecurityProfile", string(data)}, " ")
}
