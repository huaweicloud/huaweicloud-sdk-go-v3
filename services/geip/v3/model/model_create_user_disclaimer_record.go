package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUserDisclaimerRecord struct {

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o CreateUserDisclaimerRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserDisclaimerRecord struct{}"
	}

	return strings.Join([]string{"CreateUserDisclaimerRecord", string(data)}, " ")
}
