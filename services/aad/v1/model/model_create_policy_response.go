package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyResponse Response Object
type CreatePolicyResponse struct {

	// 策略id
	Id *string `json:"id,omitempty"`

	// 策略名
	Name *string `json:"name,omitempty"`

	// 实例id
	PackageId *string `json:"package_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 清洗阈值
	CleanThreshold *int32 `json:"clean_threshold,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyResponse", string(data)}, " ")
}
