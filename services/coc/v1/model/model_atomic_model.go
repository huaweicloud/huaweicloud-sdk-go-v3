package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AtomicModel struct {

	// 原子能力唯一标识
	AtomicUniqueKey *string `json:"atomic_unique_key,omitempty"`

	// 中文名
	AtomicNameZh *string `json:"atomic_name_zh,omitempty"`

	// 英文名
	AtomicNameEn *string `json:"atomic_name_en,omitempty"`

	// 标签信息：CLOUD_API、FUNCTION
	Tags *[]string `json:"tags,omitempty"`
}

func (o AtomicModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtomicModel struct{}"
	}

	return strings.Join([]string{"AtomicModel", string(data)}, " ")
}
