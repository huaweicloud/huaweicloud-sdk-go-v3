package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServerNicsReq This is a auto create Body Object
type AddServerNicsReq struct {

	//
	Nics []ServerNicsReq `json:"nics"`
}

func (o AddServerNicsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServerNicsReq struct{}"
	}

	return strings.Join([]string{"AddServerNicsReq", string(data)}, " ")
}
