package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPresetLabelResponse Response Object
type ListPresetLabelResponse struct {

	// 预置标签数量
	Count *int64 `json:"count,omitempty"`

	// 预置标签
	Labels         *[]PresetLabelRsp `json:"labels,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListPresetLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPresetLabelResponse struct{}"
	}

	return strings.Join([]string{"ListPresetLabelResponse", string(data)}, " ")
}
