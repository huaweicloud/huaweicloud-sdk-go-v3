package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateUsersReq struct {

	// DDM实例帐号相关信息的集合。
	Users []CreateUsersInfo `json:"users" xml:"users"`
}

func (o CreateUsersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUsersReq struct{}"
	}

	return strings.Join([]string{"CreateUsersReq", string(data)}, " ")
}
