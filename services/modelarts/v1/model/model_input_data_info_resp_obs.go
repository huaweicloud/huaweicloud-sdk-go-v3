package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputDataInfoRespObs **参数解释**：数据输入输出信息为OBS方式。
type InputDataInfoRespObs struct {

	// **参数解释**：训练作业需要的数据集OBS路径URL。如：“/usr/data/”。 **取值范围**：不涉及。
	ObsUrl string `json:"obs_url"`
}

func (o InputDataInfoRespObs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputDataInfoRespObs struct{}"
	}

	return strings.Join([]string{"InputDataInfoRespObs", string(data)}, " ")
}
