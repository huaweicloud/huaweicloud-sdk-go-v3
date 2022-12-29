package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisassociaterouterReq struct {
	Router *Router `json:"router"`
}

func (o DisassociaterouterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociaterouterReq struct{}"
	}

	return strings.Join([]string{"DisassociaterouterReq", string(data)}, " ")
}
