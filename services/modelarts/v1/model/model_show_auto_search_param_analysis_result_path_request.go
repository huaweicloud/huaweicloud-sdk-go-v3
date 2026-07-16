package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchParamAnalysisResultPathRequest Request Object
type ShowAutoSearchParamAnalysisResultPathRequest struct {

	// 搜索参数名称
	ParameterName string `json:"parameter_name"`

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`
}

func (o ShowAutoSearchParamAnalysisResultPathRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchParamAnalysisResultPathRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchParamAnalysisResultPathRequest", string(data)}, " ")
}
