package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserReq This is a auto update request Object
type UpdateUserReq struct {
	User *UpdateUserDetailReq `json:"user"`
}

func (o UpdateUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserReq struct{}"
	}

	return strings.Join([]string{"UpdateUserReq", string(data)}, " ")
}
