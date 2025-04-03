package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAgencyCredentialsRequest Request Object
type GetAgencyCredentialsRequest struct {

	// 创建令牌接口调用签发的访问令牌
	AccessToken string `json:"access-token"`

	// 目标账户的全局唯一标识符（ID）
	TargetAccountId string `json:"target_account_id"`

	// 委托或信任委托的统一资源名称（URN）
	AgencyUrn string `json:"agency_urn"`
}

func (o GetAgencyCredentialsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAgencyCredentialsRequest struct{}"
	}

	return strings.Join([]string{"GetAgencyCredentialsRequest", string(data)}, " ")
}
