package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectSuccessRateRequest Request Object
type ShowProjectSuccessRateRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 部署应用开始时间范围的左边界（包含），格式yyyy-MM-dd
	StartDate string `json:"start_date"`

	// 部署应用开始时间范围的右边界（包含），格式yyyy-MM-dd 。最大时间范围为1年。
	EndDate string `json:"end_date"`
}

func (o ShowProjectSuccessRateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectSuccessRateRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectSuccessRateRequest", string(data)}, " ")
}
