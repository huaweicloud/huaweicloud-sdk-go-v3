package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoutetableAssociateReqbody
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
