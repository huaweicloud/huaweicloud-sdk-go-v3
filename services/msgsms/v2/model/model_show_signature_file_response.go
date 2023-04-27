package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSignatureFileResponse struct {

	// 文件ID
	FileId *string `json:"file_id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件引用
	FileRef *int32 `json:"file_ref,omitempty"`

	// 文件大小
	FileSize *int32 `json:"file_size,omitempty"`

	// 文件类型
	FileType *int32 `json:"file_type,omitempty"`

	// 模块类型
	ModuleType *int32 `json:"module_type,omitempty"`

	// 操作人
	Operator *string `json:"operator,omitempty"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSignatureFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSignatureFileResponse struct{}"
	}

	return strings.Join([]string{"ShowSignatureFileResponse", string(data)}, " ")
}
