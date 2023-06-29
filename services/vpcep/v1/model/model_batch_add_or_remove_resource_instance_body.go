package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddOrRemoveResourceInstanceBody 批量添加或删除资源标签接口请求结构体
type BatchAddOrRemoveResourceInstanceBody struct {

	// 标签列表，没有标签默认为空数组。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 操作标识：仅限于 create（创建） delete（删除）
	Action string `json:"action"`
}

func (o BatchAddOrRemoveResourceInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceBody struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceBody", string(data)}, " ")
}
