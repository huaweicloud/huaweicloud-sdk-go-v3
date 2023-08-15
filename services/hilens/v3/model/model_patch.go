package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Patch 增量更新的部署参数
type Patch struct {

	// 环境变量更新
	Envs *[]Env `json:"envs,omitempty"`
}

func (o Patch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Patch struct{}"
	}

	return strings.Join([]string{"Patch", string(data)}, " ")
}
