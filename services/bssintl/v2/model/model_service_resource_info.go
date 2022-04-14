package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceResourceInfo struct {
	BasicInfo *ResourceBasicInfo `json:"basic_info,omitempty"`
}

func (o ServiceResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceResourceInfo struct{}"
	}

	return strings.Join([]string{"ServiceResourceInfo", string(data)}, " ")
}
