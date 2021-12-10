package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输出为webhook类型时的配置信息
type TaskOutputWebhook struct {
	// 结果回调地址

	Url string `json:"url"`
	// 结果回调时需要携带的请求头

	Headers *interface{} `json:"headers"`
	// 作业输出数据类别的列表，当输出类型下有这个列表时，表示希望这个输出类型下存放dataCategory列表内的数据，部分服务需要

	DataCategory *[]string `json:"data_category,omitempty"`
}

func (o TaskOutputWebhook) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputWebhook struct{}"
	}

	return strings.Join([]string{"TaskOutputWebhook", string(data)}, " ")
}
