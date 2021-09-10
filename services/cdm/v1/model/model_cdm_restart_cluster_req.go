package model

import (
	"encoding/json"

	"strings"
)

type CdmRestartClusterReq struct {
	Restart *CdmRestartClusterReqRestart `json:"restart"`
}

func (o CdmRestartClusterReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CdmRestartClusterReq struct{}"
	}

	return strings.Join([]string{"CdmRestartClusterReq", string(data)}, " ")
}
