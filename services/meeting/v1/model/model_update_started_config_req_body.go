package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会中修改配置项
type UpdateStartedConfigReqBody struct {

	// 锁定共享标志位 0:不锁定  1:锁定
	LockSharing *int32 `json:"lockSharing,omitempty" xml:"lockSharing"`

	// 允许呼入的范围 0：所有用户  2：企业内用户  3：被邀请用户
	CallInRestriction *int32 `json:"callInRestriction,omitempty" xml:"callInRestriction"`
}

func (o UpdateStartedConfigReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStartedConfigReqBody struct{}"
	}

	return strings.Join([]string{"UpdateStartedConfigReqBody", string(data)}, " ")
}
