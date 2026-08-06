package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupportNetWorkTypeResponse struct {

	// 网络类型
	NetWork *string `json:"net_work,omitempty"`

	// 引擎类型
	EngineTypes *[]string `json:"engine_types,omitempty"`
}

func (o SupportNetWorkTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportNetWorkTypeResponse struct{}"
	}

	return strings.Join([]string{"SupportNetWorkTypeResponse", string(data)}, " ")
}
