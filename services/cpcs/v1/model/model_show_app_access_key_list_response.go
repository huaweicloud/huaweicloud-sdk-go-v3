package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppAccessKeyListResponse Response Object
type ShowAppAccessKeyListResponse struct {

	// 访问密钥列表
	Result *[]AccessKeyInfo `json:"result,omitempty"`

	// 满足条件的密钥总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAppAccessKeyListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppAccessKeyListResponse struct{}"
	}

	return strings.Join([]string{"ShowAppAccessKeyListResponse", string(data)}, " ")
}
