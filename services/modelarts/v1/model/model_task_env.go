package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskEnv 精调训练环境变量信息
type TaskEnv struct {

	// 精调训练环境变量信息
	Envs *[]EnvVar `json:"envs,omitempty"`
}

func (o TaskEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskEnv struct{}"
	}

	return strings.Join([]string{"TaskEnv", string(data)}, " ")
}
