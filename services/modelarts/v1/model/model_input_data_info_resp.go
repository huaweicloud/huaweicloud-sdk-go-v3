package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputDataInfoResp 数据实际输入信息。枚举值：   - dataset：指定输入为数据集；   - obs：指定输入为OBS路径。
type InputDataInfoResp struct {
	Dataset *InputDataInfoRespDataset `json:"dataset,omitempty"`

	Obs *InputDataInfoRespObs `json:"obs,omitempty"`
}

func (o InputDataInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputDataInfoResp struct{}"
	}

	return strings.Join([]string{"InputDataInfoResp", string(data)}, " ")
}
