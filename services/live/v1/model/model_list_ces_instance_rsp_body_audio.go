package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesInstanceRspBodyAudio 音频
type ListCesInstanceRspBodyAudio struct {

	// 频道id
	Id *string `json:"id,omitempty"`

	// 频道name
	Name *string `json:"name,omitempty"`
}

func (o ListCesInstanceRspBodyAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceRspBodyAudio struct{}"
	}

	return strings.Join([]string{"ListCesInstanceRspBodyAudio", string(data)}, " ")
}
