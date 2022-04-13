package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RoutetableAssociateReqbody struct {
	Routetable *AsscoiateReq `json:"routetable"`
}

func (o RoutetableAssociateReqbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutetableAssociateReqbody struct{}"
	}

	return strings.Join([]string{"RoutetableAssociateReqbody", string(data)}, " ")
}
