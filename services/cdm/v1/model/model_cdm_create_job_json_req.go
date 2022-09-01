package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmCreateJobJsonReq struct {

	// 作业列表，请参见jobs数据结构说明。
	Jobs []Job `json:"jobs" xml:"jobs"`
}

func (o CdmCreateJobJsonReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmCreateJobJsonReq struct{}"
	}

	return strings.Join([]string{"CdmCreateJobJsonReq", string(data)}, " ")
}
