package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCesInstanceRspBodyInstances struct {
	MedialiveMpc *ListCesInstanceRspBodyMedialiveMpc `json:"medialive_mpc"`

	Pipeline *ListCesInstanceRspBodyPipeline `json:"pipeline,omitempty"`

	Audio *ListCesInstanceRspBodyAudio `json:"audio,omitempty"`
}

func (o ListCesInstanceRspBodyInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceRspBodyInstances struct{}"
	}

	return strings.Join([]string{"ListCesInstanceRspBodyInstances", string(data)}, " ")
}
