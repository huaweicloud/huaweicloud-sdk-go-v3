package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectorConfigModel 采集器配置。
type CollectorConfigModel struct {
	LineViewConfig *LineViewConfigModel `json:"line_view_config,omitempty"`

	DetailViewConfig *DetailViewConfigModel `json:"detail_view_config,omitempty"`
}

func (o CollectorConfigModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectorConfigModel struct{}"
	}

	return strings.Join([]string{"CollectorConfigModel", string(data)}, " ")
}
