package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAppGroupReq 校验应用组的请求体。
type CheckAppGroupReq struct {

	// 应用组名称。
	Name string `json:"name"`
}

func (o CheckAppGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAppGroupReq struct{}"
	}

	return strings.Join([]string{"CheckAppGroupReq", string(data)}, " ")
}
