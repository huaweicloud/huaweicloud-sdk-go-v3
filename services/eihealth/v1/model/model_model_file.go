package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelFile 模型数据文件
type ModelFile struct {
	Source *ModelFileSource `json:"source"`

	// 文件URL，用户私有数据中心为项目路径、公共数据场景为obs地址
	Url string `json:"url"`

	// 模型文件所在项目id，仅文件为数据中心时填写
	EihealthProjectId *string `json:"eihealth_project_id,omitempty"`
}

func (o ModelFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelFile struct{}"
	}

	return strings.Join([]string{"ModelFile", string(data)}, " ")
}
