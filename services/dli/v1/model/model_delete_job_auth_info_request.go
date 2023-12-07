package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobAuthInfoRequest Request Object
type DeleteJobAuthInfoRequest struct {

	// 即将删除的认证信息名
	AuthInfoName string `json:"auth_info_name"`
}

func (o DeleteJobAuthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobAuthInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobAuthInfoRequest", string(data)}, " ")
}
