package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyEnterpriseIdReq 修改企业ID请求
type ModifyEnterpriseIdReq struct {

	// 企业ID
	EnterpriseId string `json:"enterprise_id"`
}

func (o ModifyEnterpriseIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEnterpriseIdReq struct{}"
	}

	return strings.Join([]string{"ModifyEnterpriseIdReq", string(data)}, " ")
}
