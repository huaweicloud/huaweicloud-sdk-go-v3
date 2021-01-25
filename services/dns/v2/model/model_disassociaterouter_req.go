package model

import (
	"encoding/json"

	"strings"
)

// 需要解关联的Router(VPC)。
type DisassociaterouterReq struct {
	Router *Router `json:"router"`
}

func (o DisassociaterouterReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociaterouterReq struct{}"
	}

	return strings.Join([]string{"DisassociaterouterReq", string(data)}, " ")
}
