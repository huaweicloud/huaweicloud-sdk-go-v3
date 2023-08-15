package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestAllowClientRecordReqBody 允许客户端录制请求。
type RestAllowClientRecordReqBody struct {

	// * 0：取消与会者客户端录制权限 * 1：允许与会者客户端录制
	AllowClientRecord int32 `json:"allowClientRecord"`
}

func (o RestAllowClientRecordReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestAllowClientRecordReqBody struct{}"
	}

	return strings.Join([]string{"RestAllowClientRecordReqBody", string(data)}, " ")
}
