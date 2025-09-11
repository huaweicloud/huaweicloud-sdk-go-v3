package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollupEnableResp **参数解释** 是否开启聚合 **取值范围** - true：表示开启聚合 - false：表示不开启聚合
type RollupEnableResp struct {
}

func (o RollupEnableResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollupEnableResp struct{}"
	}

	return strings.Join([]string{"RollupEnableResp", string(data)}, " ")
}
