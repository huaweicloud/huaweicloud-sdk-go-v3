package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataLayerVoList struct {

	// 数仓分层数组。
	Levels *[]DataLayerVo `json:"levels,omitempty"`
}

func (o DataLayerVoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataLayerVoList struct{}"
	}

	return strings.Join([]string{"DataLayerVoList", string(data)}, " ")
}
