package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDependencyResponse Response Object
type CreateDependencyResponse struct {

	// 依赖包版本ID。
	Id *string `json:"id,omitempty"`

	// 依赖包拥有者。
	Owner *string `json:"owner,omitempty"`

	// 依赖包在obs的存储地址。
	Link *string `json:"link,omitempty"`

	// 运行时语言。
	Runtime *string `json:"runtime,omitempty"`

	// 依赖包唯一标志。
	Etag *string `json:"etag,omitempty"`

	// 依赖包大小。
	Size *int64 `json:"size,omitempty"`

	// 依赖包名。
	Name *string `json:"name,omitempty"`

	// 依赖包描述。
	Description *string `json:"description,omitempty"`

	// 依赖包文件名。
	FileName *string `json:"file_name,omitempty"`

	// 依赖包版本号
	Version *int64 `json:"version,omitempty"`

	// 依赖包ID
	DepId *string `json:"dep_id,omitempty"`

	// 依赖包更新时间
	LastModified   *int64 `json:"last_modified,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateDependencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDependencyResponse struct{}"
	}

	return strings.Join([]string{"CreateDependencyResponse", string(data)}, " ")
}
