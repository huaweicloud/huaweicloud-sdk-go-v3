package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkFlowModel struct {

	// 英文描述
	EnUs map[string]string `json:"en-us,omitempty"`

	// 中文描述
	ZhCn map[string]string `json:"zh-cn,omitempty"`
}

func (o WorkFlowModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkFlowModel struct{}"
	}

	return strings.Join([]string{"WorkFlowModel", string(data)}, " ")
}
