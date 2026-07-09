package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealNameAuthStatusResponse Response Object
type ShowRealNameAuthStatusResponse struct {

	// 实名认证状态。enum:-1,0,1,2。 -1未实名认证、0实名认证审核中、1实名认证不通过、2已实名认证
	VerifiedStatus *int32 `json:"verified_status,omitempty"`

	// 实名认证类型。实名认证状态为-1未实名认证返回null。enum:0,1。 0个人实名认证、1企业实名认证
	VerifiedType   *int32 `json:"verified_type,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRealNameAuthStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealNameAuthStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowRealNameAuthStatusResponse", string(data)}, " ")
}
