package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckRejoinDomainResult 加域检查桌面信息。
type BatchCheckRejoinDomainResult struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 镜像名称。
	ImageName *string `json:"image_name,omitempty"`

	// 加域状态。|- 1 正常。 2 脱域。 3 未上报。
	DomainStatus *int32 `json:"domain_status,omitempty"`

	// 是否可以加域。
	RejoinAble *bool `json:"rejoin_able,omitempty"`

	Product *ProductInfo `json:"product,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchCheckRejoinDomainResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckRejoinDomainResult struct{}"
	}

	return strings.Join([]string{"BatchCheckRejoinDomainResult", string(data)}, " ")
}
