package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestScenePageParam struct {
	Params *ScenePageParam `json:"params,omitempty"`
}

func (o CommRequestScenePageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestScenePageParam struct{}"
	}

	return strings.Join([]string{"CommRequestScenePageParam", string(data)}, " ")
}
