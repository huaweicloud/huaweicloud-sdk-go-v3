package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateSeverGroupReq 校验服务器组的请求体。
type ValidateSeverGroupReq struct {

	// 服务器组名称。
	Name string `json:"name"`
}

func (o ValidateSeverGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateSeverGroupReq struct{}"
	}

	return strings.Join([]string{"ValidateSeverGroupReq", string(data)}, " ")
}
