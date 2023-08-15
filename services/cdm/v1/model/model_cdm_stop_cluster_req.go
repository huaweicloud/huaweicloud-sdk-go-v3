package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmStopClusterReq struct {
	Stop *CdmStopClusterReqStop `json:"stop"`
}

func (o CdmStopClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmStopClusterReq struct{}"
	}

	return strings.Join([]string{"CdmStopClusterReq", string(data)}, " ")
}
