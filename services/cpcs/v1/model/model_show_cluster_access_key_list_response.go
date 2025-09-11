package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterAccessKeyListResponse Response Object
type ShowClusterAccessKeyListResponse struct {

	// 已授权的访问密钥列表
	Result         *[]ShowClusterAccessKeyListResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                           `json:"-"`
}

func (o ShowClusterAccessKeyListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterAccessKeyListResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterAccessKeyListResponse", string(data)}, " ")
}
