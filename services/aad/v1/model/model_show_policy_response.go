package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyResponse Response Object
type ShowPolicyResponse struct {

	// 策略id
	Id *string `json:"id,omitempty"`

	// 防护包id
	PackageId *string `json:"package_id,omitempty"`

	// 策略名
	Name *string `json:"name,omitempty"`

	// 清洗阈值
	CleanThreshold *int32 `json:"clean_threshold,omitempty"`

	PopPolicy      *PopPolicy `json:"pop_policy,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyResponse", string(data)}, " ")
}
