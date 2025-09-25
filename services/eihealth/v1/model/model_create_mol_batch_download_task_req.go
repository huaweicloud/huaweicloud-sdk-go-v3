package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMolBatchDownloadTaskReq 创建分子或分子复合物批量下载任务请求体
type CreateMolBatchDownloadTaskReq struct {

	// 作业ID
	JobId string `json:"job_id"`

	// 作业结果文件url
	JobResultUrl string `json:"job_result_url"`

	// 下载类型：MOL（小分子）、COMPLEX（复合物）
	Mode string `json:"mode"`

	// **参数解释**： 文件格式：pdb、sdf（仅下载小分子时支持此格式）。 **约束限制**： 不涉及 **取值范围**： - pdb：pdb文件格式。 - sdf：sdf文件格式，仅下载小分子时支持此格式。 **默认取值**： pdb
	Format *string `json:"format,omitempty"`

	// 选中下载的分子下标
	Selected []int32 `json:"selected"`
}

func (o CreateMolBatchDownloadTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMolBatchDownloadTaskReq struct{}"
	}

	return strings.Join([]string{"CreateMolBatchDownloadTaskReq", string(data)}, " ")
}
