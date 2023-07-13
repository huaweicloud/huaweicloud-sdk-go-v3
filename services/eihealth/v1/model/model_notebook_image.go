package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotebookImage notebook镜像
type NotebookImage struct {
	ImageType *DevelopImageType `json:"image_type"`

	ImageInfo *ImageInfo `json:"image_info"`
}

func (o NotebookImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookImage struct{}"
	}

	return strings.Join([]string{"NotebookImage", string(data)}, " ")
}
