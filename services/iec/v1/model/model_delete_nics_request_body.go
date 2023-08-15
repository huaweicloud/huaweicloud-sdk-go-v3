package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNicsRequestBody 删除网卡请求体。
type DeleteNicsRequestBody struct {

	// 待删除网卡信息。
	Nics []NicId `json:"nics"`
}

func (o DeleteNicsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNicsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteNicsRequestBody", string(data)}, " ")
}
