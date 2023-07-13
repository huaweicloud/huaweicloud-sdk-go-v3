package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeShareNameReq change_name对象
type ChangeShareNameReq struct {
	ChangeName *ShareName `json:"change_name"`
}

func (o ChangeShareNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareNameReq struct{}"
	}

	return strings.Join([]string{"ChangeShareNameReq", string(data)}, " ")
}
