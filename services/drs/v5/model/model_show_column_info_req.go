package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowColumnInfoReq 查询列信息请求体
type ShowColumnInfoReq struct {

	// 是否重新从node获取列信息
	Refresh bool `json:"refresh"`

	// 列所属的对象信息
	ObjectIds *[]string `json:"object_ids,omitempty"`
}

func (o ShowColumnInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowColumnInfoReq struct{}"
	}

	return strings.Join([]string{"ShowColumnInfoReq", string(data)}, " ")
}
