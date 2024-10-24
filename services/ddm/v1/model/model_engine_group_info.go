package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineGroupInfo struct {

	// 引擎id。
	Id string `json:"id"`

	// 引擎名称。
	Name string `json:"name"`

	// 引擎版本。
	Version string `json:"version"`
}

func (o EngineGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineGroupInfo struct{}"
	}

	return strings.Join([]string{"EngineGroupInfo", string(data)}, " ")
}
