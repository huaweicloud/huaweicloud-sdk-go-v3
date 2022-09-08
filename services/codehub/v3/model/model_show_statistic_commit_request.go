package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowStatisticCommitRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`

	// 分支名称
	RefName string `json:"ref_name"`

	// 起始提交日期,格式为yyyy-MM-dd
	BeginDate string `json:"begin_date"`

	// 终止提交日期,格式为yyyy-MM-dd（begin_date和end_date时间间隔不超过60天）
	EndDate string `json:"end_date"`
}

func (o ShowStatisticCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticCommitRequest struct{}"
	}

	return strings.Join([]string{"ShowStatisticCommitRequest", string(data)}, " ")
}
