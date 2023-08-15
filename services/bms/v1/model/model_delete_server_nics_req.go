package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerNicsReq 需要解绑的网卡列表信息
type DeleteServerNicsReq struct {

	//
	Nics []ServerNics `json:"nics"`
}

func (o DeleteServerNicsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerNicsReq struct{}"
	}

	return strings.Join([]string{"DeleteServerNicsReq", string(data)}, " ")
}
