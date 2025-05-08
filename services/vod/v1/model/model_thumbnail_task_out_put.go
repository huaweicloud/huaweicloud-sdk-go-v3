package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThumbnailTaskOutPut struct {
	ObsInfo *ObsInfo `json:"obs_info,omitempty"`

	// 截图信息列表
	PicInfoList *[]PicInfo `json:"pic_info_list,omitempty"`
}

func (o ThumbnailTaskOutPut) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThumbnailTaskOutPut struct{}"
	}

	return strings.Join([]string{"ThumbnailTaskOutPut", string(data)}, " ")
}
