package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPresetLabelRequest Request Object
type ListPresetLabelRequest struct {
}

func (o ListPresetLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPresetLabelRequest struct{}"
	}

	return strings.Join([]string{"ListPresetLabelRequest", string(data)}, " ")
}
