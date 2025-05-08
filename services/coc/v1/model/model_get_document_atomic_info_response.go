package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDocumentAtomicInfoResponse Response Object
type GetDocumentAtomicInfoResponse struct {

	// 原子能力唯一标识：只允许字母+下划线，字母开头
	AtomicUniqueKey *string `json:"atomic_unique_key,omitempty"`

	// 中文名
	AtomicNameZh *string `json:"atomic_name_zh,omitempty"`

	// 英文名
	AtomicNameEn *string `json:"atomic_name_en,omitempty"`

	// 标签信息
	Tags *[]string `json:"tags,omitempty"`

	// 原子能力入参
	Inputs *[]AtomicInputModel `json:"inputs,omitempty"`

	Outputs        *AtomicOutputModel `json:"outputs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o GetDocumentAtomicInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDocumentAtomicInfoResponse struct{}"
	}

	return strings.Join([]string{"GetDocumentAtomicInfoResponse", string(data)}, " ")
}
