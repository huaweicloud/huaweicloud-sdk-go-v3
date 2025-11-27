package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Artifact struct {

	// 制品的相对文件路径
	Path *string `json:"path,omitempty"`

	// HTTP地址，可通过该地址下载或访问制品内容
	Url *string `json:"url,omitempty"`

	// 版本标识符
	Revision *string `json:"revision,omitempty"`

	// 文件摘要，格式为 <算法>:<校验值>
	Digest *string `json:"digest,omitempty"`

	// 最后更新时间
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 文件大小（以字节为单位）
	Size *int32 `json:"size,omitempty"`
}

func (o Artifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Artifact struct{}"
	}

	return strings.Join([]string{"Artifact", string(data)}, " ")
}
