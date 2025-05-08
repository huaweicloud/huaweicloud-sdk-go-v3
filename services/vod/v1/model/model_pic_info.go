package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PicInfo struct {

	// 截图在obs桶中路径
	ObjectKey *string `json:"object_key,omitempty"`
}

func (o PicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PicInfo struct{}"
	}

	return strings.Join([]string{"PicInfo", string(data)}, " ")
}
