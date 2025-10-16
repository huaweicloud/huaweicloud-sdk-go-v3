package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorizeTxtRecordVerificationResponse Response Object
type CreateAuthorizeTxtRecordVerificationResponse struct {

	// **参数解释：** 授权状态。 **取值范围：** - CREATED：已创建 - VERIFIED：已验证
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuthorizeTxtRecordVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizeTxtRecordVerificationResponse struct{}"
	}

	return strings.Join([]string{"CreateAuthorizeTxtRecordVerificationResponse", string(data)}, " ")
}
