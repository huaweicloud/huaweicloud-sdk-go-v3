package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedAtlas struct {

	// 图库ID
	AtlasId *string `json:"atlas_id,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`
}

func (o FailedAtlas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedAtlas struct{}"
	}

	return strings.Join([]string{"FailedAtlas", string(data)}, " ")
}
