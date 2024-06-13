package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PopUpInfo struct {

	// 是否弹窗
	Popup *bool `json:"popup,omitempty"`

	// 包周期计费时长上限
	TimeLimit *interface{} `json:"time_limit,omitempty"`
}

func (o PopUpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PopUpInfo struct{}"
	}

	return strings.Join([]string{"PopUpInfo", string(data)}, " ")
}
