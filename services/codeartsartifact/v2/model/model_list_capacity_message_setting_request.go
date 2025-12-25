package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCapacityMessageSettingRequest Request Object
type ListCapacityMessageSettingRequest struct {
}

func (o ListCapacityMessageSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCapacityMessageSettingRequest struct{}"
	}

	return strings.Join([]string{"ListCapacityMessageSettingRequest", string(data)}, " ")
}
