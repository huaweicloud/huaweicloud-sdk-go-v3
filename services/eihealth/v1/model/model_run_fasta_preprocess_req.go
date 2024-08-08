package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFastaPreprocessReq Fasta受体预览请求体
type RunFastaPreprocessReq struct {
	File *FastaReceptor `json:"file"`

	// 预览数量
	PreviewCount *int32 `json:"preview_count,omitempty"`

	// 计数上限
	CountLimit *int32 `json:"count_limit,omitempty"`
}

func (o RunFastaPreprocessReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFastaPreprocessReq struct{}"
	}

	return strings.Join([]string{"RunFastaPreprocessReq", string(data)}, " ")
}
