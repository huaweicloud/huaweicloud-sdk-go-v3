package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签信息
type UpadteNodeLabelReq struct {

	// 标签名称
	Name string `json:"name"`
}

func (o UpadteNodeLabelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpadteNodeLabelReq struct{}"
	}

	return strings.Join([]string{"UpadteNodeLabelReq", string(data)}, " ")
}
