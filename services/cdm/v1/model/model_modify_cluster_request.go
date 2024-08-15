package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyClusterRequest Request Object
type ModifyClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 请求语言。
	XLanguage string `json:"X-Language"`

	Body *CdmModifyClusterReq `json:"body,omitempty"`
}

func (o ModifyClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyClusterRequest struct{}"
	}

	return strings.Join([]string{"ModifyClusterRequest", string(data)}, " ")
}
