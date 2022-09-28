package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 导入镜像UI接口请求体
type ImportImageReq struct {

	// 源项目ID
	SourceProjectId string `json:"source_project_id"`

	// 镜像ID
	ImageId string `json:"image_id"`

	// 镜像tag
	Tag string `json:"tag"`
}

func (o ImportImageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageReq struct{}"
	}

	return strings.Join([]string{"ImportImageReq", string(data)}, " ")
}
