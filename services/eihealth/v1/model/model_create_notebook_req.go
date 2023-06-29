package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotebookReq 创建notebook请求体
type CreateNotebookReq struct {

	// 描述信息，取值范围[0,1024]
	Description *string `json:"description,omitempty"`

	// 挂载信息
	Storages []NotebookStorage `json:"storages"`

	Flavor *FlavorInfo `json:"flavor"`

	Image *NotebookImage `json:"image"`

	// notebook名称，取值范围[1,63],仅支持小写字母、数字、中划线(-),开始只能是小写字母，结束只能是小写字母或数字
	Name string `json:"name"`
}

func (o CreateNotebookReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotebookReq struct{}"
	}

	return strings.Join([]string{"CreateNotebookReq", string(data)}, " ")
}
