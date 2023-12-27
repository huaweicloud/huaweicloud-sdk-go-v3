package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdeTrashArtifactModel struct {

	// 仓库id
	Id *string `json:"id,omitempty"`

	// 类型
	Format *string `json:"format,omitempty"`

	// 当前仓库状态
	Status *string `json:"status,omitempty"`

	// 待还原的文件路径
	Uri *string `json:"uri,omitempty"`
}

func (o IdeTrashArtifactModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdeTrashArtifactModel struct{}"
	}

	return strings.Join([]string{"IdeTrashArtifactModel", string(data)}, " ")
}
