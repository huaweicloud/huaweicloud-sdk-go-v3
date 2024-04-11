package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsResultData 返回的数据信息。
type TagsResultData struct {

	// 给表/属性打标签/删除标签的接口返回，返回的vale是null则表示成功。
	Value *string `json:"value,omitempty"`
}

func (o TagsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsResultData struct{}"
	}

	return strings.Join([]string{"TagsResultData", string(data)}, " ")
}
