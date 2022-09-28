package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新notebook请求体
type UpdateNotebookReq struct {

	// notebook描述信息，取值范围[0,1024]
	Description string `json:"description"`
}

func (o UpdateNotebookReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotebookReq struct{}"
	}

	return strings.Join([]string{"UpdateNotebookReq", string(data)}, " ")
}
