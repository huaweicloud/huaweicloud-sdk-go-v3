package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunCheckPictureResponse struct {
	// 调用成功时表示调用结果。  调用失败时无此字段。  - true表示图像索引库中存在查询的图片。  - false表示图像索引库中不存在查询的图片。

	Exist          *string `json:"exist,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunCheckPictureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCheckPictureResponse struct{}"
	}

	return strings.Join([]string{"RunCheckPictureResponse", string(data)}, " ")
}
