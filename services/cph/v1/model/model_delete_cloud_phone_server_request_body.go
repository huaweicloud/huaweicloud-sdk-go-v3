package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudPhoneServerRequestBody 删除云手机服务器请求体。
type DeleteCloudPhoneServerRequestBody struct {

	// 云手机服务器id列表。
	ServerIds []string `json:"server_ids"`
}

func (o DeleteCloudPhoneServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudPhoneServerRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteCloudPhoneServerRequestBody", string(data)}, " ")
}
