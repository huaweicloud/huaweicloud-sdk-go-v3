package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmRestartClusterReq struct {
	Restart *CdmRestartClusterReqRestart `json:"restart"`
}

func (o CdmRestartClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmRestartClusterReq struct{}"
	}

	return strings.Join([]string{"CdmRestartClusterReq", string(data)}, " ")
}
