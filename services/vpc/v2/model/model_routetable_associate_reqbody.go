package model

import (
	"encoding/json"

	"strings"
)

//
type RoutetableAssociateReqbody struct {
	Routetable *AsscoiateReq `json:"routetable"`
}

func (o RoutetableAssociateReqbody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RoutetableAssociateReqbody struct{}"
	}

	return strings.Join([]string{"RoutetableAssociateReqbody", string(data)}, " ")
}
