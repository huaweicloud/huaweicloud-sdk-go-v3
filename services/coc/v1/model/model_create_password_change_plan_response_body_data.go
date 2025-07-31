package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePasswordChangePlanResponseBodyData 改密计划具体数据
type CreatePasswordChangePlanResponseBodyData struct {

	// 改密计划总条数
	Total int32 `json:"total"`

	// 改密计划具体信息
	List []AccountChangePwdPlan `json:"list"`
}

func (o CreatePasswordChangePlanResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePasswordChangePlanResponseBodyData struct{}"
	}

	return strings.Join([]string{"CreatePasswordChangePlanResponseBodyData", string(data)}, " ")
}
