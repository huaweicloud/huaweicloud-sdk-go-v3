package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeGraphReqResize resize是一个对象
type ResizeGraphReqResize struct {

	// 图规格类型，当前支持取值为\"2\",\"3\",\"4\",\"5\"分别代表扩容成千万边、一亿边、十亿边、百亿边规格的图。
	GraphSizeTypeIndex string `json:"graph_size_type_index"`
}

func (o ResizeGraphReqResize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeGraphReqResize struct{}"
	}

	return strings.Join([]string{"ResizeGraphReqResize", string(data)}, " ")
}
