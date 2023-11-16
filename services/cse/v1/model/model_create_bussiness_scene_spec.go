package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBussinessSceneSpec 流量特征描述
type CreateBussinessSceneSpec struct {

	// 特征名称
	Alias *string `json:"alias,omitempty"`

	// 匹配条件定义
	Matches *[]CreateBussinessSceneSpecMatches `json:"matches,omitempty"`
}

func (o CreateBussinessSceneSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBussinessSceneSpec struct{}"
	}

	return strings.Join([]string{"CreateBussinessSceneSpec", string(data)}, " ")
}
