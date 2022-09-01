package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDependencyResponse struct {

	// 依赖包ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 依赖包拥有者。
	Owner *string `json:"owner,omitempty" xml:"owner"`

	// 依赖包在obs的存储地址。
	Link *string `json:"link,omitempty" xml:"link"`

	// 运行时语言。
	Runtime *string `json:"runtime,omitempty" xml:"runtime"`

	// 依赖包唯一标志。
	Etag *string `json:"etag,omitempty" xml:"etag"`

	// 依赖包大小。
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 依赖包名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 依赖包描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 依赖包文件名。
	FileName       *string `json:"file_name,omitempty" xml:"file_name"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDependencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDependencyResponse struct{}"
	}

	return strings.Join([]string{"CreateDependencyResponse", string(data)}, " ")
}
