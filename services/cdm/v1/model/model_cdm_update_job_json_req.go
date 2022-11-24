package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmUpdateJobJsonReq struct {

	// 作业列表，请参见jobs数据结构说明。
	Jobs []Job `json:"jobs"`
}

func (o CdmUpdateJobJsonReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmUpdateJobJsonReq struct{}"
	}

	return strings.Join([]string{"CdmUpdateJobJsonReq", string(data)}, " ")
}
