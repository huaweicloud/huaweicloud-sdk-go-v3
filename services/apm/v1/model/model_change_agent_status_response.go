package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAgentStatusResponse Response Object
type ChangeAgentStatusResponse struct {

	// 返回结果ok表示成功。
	Flag           *string `json:"flag,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeAgentStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAgentStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeAgentStatusResponse", string(data)}, " ")
}
