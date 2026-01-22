package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResumeInstanceReq struct {

	// **参数解释**： 实例列表。
	Instances *[]string `json:"instances,omitempty"`
}

func (o BatchResumeInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResumeInstanceReq struct{}"
	}

	return strings.Join([]string{"BatchResumeInstanceReq", string(data)}, " ")
}
