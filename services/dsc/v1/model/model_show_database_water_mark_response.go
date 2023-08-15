package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseWaterMarkResponse Response Object
type ShowDatabaseWaterMarkResponse struct {

	// 提取水印内容列表。上传数据中不同列可能包含不同水印，返回时将所有提取到的水印返回，列表中水印个数不超过100
	Watermarks     *[]string `json:"watermarks,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowDatabaseWaterMarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseWaterMarkResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseWaterMarkResponse", string(data)}, " ")
}
