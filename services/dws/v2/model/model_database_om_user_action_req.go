package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseOmUserActionReq 数据库运维用户操作请求
type DatabaseOmUserActionReq struct {

	// 操作类型，取值如下：  - addOmUser：添加运维用户 - deleteOmUser：删除运维用户 - increaseOmUserPeriod ：延长用户有效期
	Operation string `json:"operation"`
}

func (o DatabaseOmUserActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseOmUserActionReq struct{}"
	}

	return strings.Join([]string{"DatabaseOmUserActionReq", string(data)}, " ")
}
