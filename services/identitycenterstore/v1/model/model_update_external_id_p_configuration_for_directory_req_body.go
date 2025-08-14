package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateExternalIdPConfigurationForDirectoryReqBody struct {

	// 身份提供商发布者标识
	EntityId string `json:"entity_id"`

	// 身份提供商登录链接
	LoginUrl string `json:"login_url"`
}

func (o UpdateExternalIdPConfigurationForDirectoryReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExternalIdPConfigurationForDirectoryReqBody struct{}"
	}

	return strings.Join([]string{"UpdateExternalIdPConfigurationForDirectoryReqBody", string(data)}, " ")
}
