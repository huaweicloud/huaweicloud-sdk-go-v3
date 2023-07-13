package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDiagnoseItemsReq struct {

	// job id
	JobId *string `json:"job_id,omitempty"`

	// 诊断项id列表
	ItemIds *string `json:"item_ids,omitempty"`
}

func (o QueryDiagnoseItemsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDiagnoseItemsReq struct{}"
	}

	return strings.Join([]string{"QueryDiagnoseItemsReq", string(data)}, " ")
}
