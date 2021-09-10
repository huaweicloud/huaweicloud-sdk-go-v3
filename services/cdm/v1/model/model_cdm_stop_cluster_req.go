package model

import (
	"encoding/json"

	"strings"
)

type CdmStopClusterReq struct {
	Stop *CdmStopClusterReqStop `json:"stop"`
}

func (o CdmStopClusterReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CdmStopClusterReq struct{}"
	}

	return strings.Join([]string{"CdmStopClusterReq", string(data)}, " ")
}
