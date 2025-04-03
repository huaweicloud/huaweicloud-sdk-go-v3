package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataServiceInstanceAccesslogsRequest Request Object
type ListDataServiceInstanceAccesslogsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType *string `json:"Content-Type,omitempty"`

	// 集群ID编号。
	InstanceId string `json:"instance_id"`

	// 是否查询API的访问日志，true表示查询API的访问日志，false表示查询应用的访问日志。
	IsApi *bool `json:"is_api,omitempty"`
}

func (o ListDataServiceInstanceAccesslogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataServiceInstanceAccesslogsRequest struct{}"
	}

	return strings.Join([]string{"ListDataServiceInstanceAccesslogsRequest", string(data)}, " ")
}
