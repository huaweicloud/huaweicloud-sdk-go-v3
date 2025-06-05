package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeystoreResponse Response Object
type ListKeystoreResponse struct {

	// 文件列表
	Result *[]ListKeystoreResult `json:"result,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListKeystoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreResponse struct{}"
	}

	return strings.Join([]string{"ListKeystoreResponse", string(data)}, " ")
}
