package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEnterpriseIdReq 检查企业ID是否已经使用的请求。
type CheckEnterpriseIdReq struct {

	// 企业ID。
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

func (o CheckEnterpriseIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEnterpriseIdReq struct{}"
	}

	return strings.Join([]string{"CheckEnterpriseIdReq", string(data)}, " ")
}
