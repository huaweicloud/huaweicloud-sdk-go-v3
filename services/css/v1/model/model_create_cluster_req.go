package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateClusterReq struct {
	Cluster *CreateClusterBody `json:"cluster"`
}

func (o CreateClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReq struct{}"
	}

	return strings.Join([]string{"CreateClusterReq", string(data)}, " ")
}
