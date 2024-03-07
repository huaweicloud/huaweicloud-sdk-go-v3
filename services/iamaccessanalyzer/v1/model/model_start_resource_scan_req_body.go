package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartResourceScanReqBody struct {

	// 资源的唯一标识符。
	ResourceId *string `json:"resource_id,omitempty"`

	// 拥有资源的账户ID。
	ResourceOwnerAccount string `json:"resource_owner_account"`

	// 访问分析的唯一资源标识。
	ResourceUrn string `json:"resource_urn"`
}

func (o StartResourceScanReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResourceScanReqBody struct{}"
	}

	return strings.Join([]string{"StartResourceScanReqBody", string(data)}, " ")
}
