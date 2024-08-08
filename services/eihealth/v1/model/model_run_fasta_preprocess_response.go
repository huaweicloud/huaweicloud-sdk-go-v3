package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFastaPreprocessResponse Response Object
type RunFastaPreprocessResponse struct {

	// 已知的蛋白序列数量
	Count *int32 `json:"count,omitempty"`

	// 文件中是否还有更多氨基酸序列没有处理
	HasMore        *bool `json:"has_more,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o RunFastaPreprocessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFastaPreprocessResponse struct{}"
	}

	return strings.Join([]string{"RunFastaPreprocessResponse", string(data)}, " ")
}
