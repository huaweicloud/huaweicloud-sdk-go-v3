package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputDataInfo 数据实际输入信息。枚举值：   - dataset：指定输入为数据集；   - obs：指定输入为OBS路径。
type InputDataInfo struct {
	Dataset *InputDataInfoDataset `json:"dataset,omitempty"`

	Obs *InputDataInfoObs `json:"obs,omitempty"`
}

func (o InputDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputDataInfo struct{}"
	}

	return strings.Join([]string{"InputDataInfo", string(data)}, " ")
}
