package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GraphSizeTypeIndexReq struct {
	// 图规格类型，当前支持取值为\"2\",\"3\",\"4\",\"5\"分别代表扩容成千万边、一亿边、十亿边、百亿边规格的图

	GraphSizeTypeIndex string `json:"graphSizeTypeIndex"`
}

func (o GraphSizeTypeIndexReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GraphSizeTypeIndexReq struct{}"
	}

	return strings.Join([]string{"GraphSizeTypeIndexReq", string(data)}, " ")
}
