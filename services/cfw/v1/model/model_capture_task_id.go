package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaptureTaskId struct {

	// 防火墙id
	Id *string `json:"id,omitempty"`

	// 防火墙名称
	Name *string `json:"name,omitempty"`
}

func (o CaptureTaskId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaptureTaskId struct{}"
	}

	return strings.Join([]string{"CaptureTaskId", string(data)}, " ")
}
