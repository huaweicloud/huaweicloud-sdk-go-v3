package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLigandPreviewTaskReq 创建配体预览任务请求体
type CreateLigandPreviewTaskReq struct {
	LigandFile *DrugFile `json:"ligand_file"`

	// 预览数量，若分子数量大于预览数量，则超出预览数量部分只做计数
	PreviewCount *int32 `json:"preview_count,omitempty"`

	// 计数上限，若分子数量大于计数上限，则终止计数并在结果中标明计数不完整（has_more=true），计数数量应不小于preview_count
	CountLimit *int32 `json:"count_limit,omitempty"`
}

func (o CreateLigandPreviewTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLigandPreviewTaskReq struct{}"
	}

	return strings.Join([]string{"CreateLigandPreviewTaskReq", string(data)}, " ")
}
