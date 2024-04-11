package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartInstanceRequestBody struct {

	// 需要启动的节点的ID，取值不能为null，如果为空列表则启动整个实例
	NodeIds []string `json:"node_ids"`
}

func (o StartInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"StartInstanceRequestBody", string(data)}, " ")
}
