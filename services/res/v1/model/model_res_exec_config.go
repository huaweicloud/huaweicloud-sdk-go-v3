package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ResExecConfig struct {
	SparkCalcSpec *SparkCalcSpec `json:"spark_calc_spec,omitempty" xml:"spark_calc_spec"`

	// spark可选配置项
	SparkOptionConfs *[]SparkOptionConf `json:"spark_option_confs,omitempty" xml:"spark_option_confs"`
}

func (o ResExecConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResExecConfig struct{}"
	}

	return strings.Join([]string{"ResExecConfig", string(data)}, " ")
}
