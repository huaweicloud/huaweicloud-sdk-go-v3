package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputDataInfoObs 数据输入输出信息为OBS方式。
type InputDataInfoObs struct {

	// 训练作业需要的数据集OBS路径URL。如：“/usr/data/”。
	ObsUrl string `json:"obs_url"`
}

func (o InputDataInfoObs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputDataInfoObs struct{}"
	}

	return strings.Join([]string{"InputDataInfoObs", string(data)}, " ")
}
