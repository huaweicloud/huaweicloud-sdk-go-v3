package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeTypeRequest Request Object
type ListVolumeTypeRequest struct {
}

func (o ListVolumeTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeTypeRequest struct{}"
	}

	return strings.Join([]string{"ListVolumeTypeRequest", string(data)}, " ")
}
