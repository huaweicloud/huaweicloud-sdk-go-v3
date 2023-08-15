package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResExecConfig
type ResExecConfig struct {
	SparkCalcSpec *SparkCalcSpec `json:"spark_calc_spec,omitempty"`

	// spark可选配置项
	SparkOptionConfs *[]SparkOptionConf `json:"spark_option_confs,omitempty"`
}

func (o ResExecConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResExecConfig struct{}"
	}

	return strings.Join([]string{"ResExecConfig", string(data)}, " ")
}
