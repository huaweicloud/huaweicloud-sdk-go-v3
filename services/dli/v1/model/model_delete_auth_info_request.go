package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuthInfoRequest Request Object
type DeleteAuthInfoRequest struct {

	// 即将删除的认证信息名
	AuthInfoName string `json:"auth_info_name"`
}

func (o DeleteAuthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuthInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuthInfoRequest", string(data)}, " ")
}
