package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeletePictureReq struct {

	// 图片URL路径，作为图片库中索引图片的ID。
	Path *string `json:"path,omitempty" xml:"path"`
}

func (o DeletePictureReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePictureReq struct{}"
	}

	return strings.Join([]string{"DeletePictureReq", string(data)}, " ")
}
