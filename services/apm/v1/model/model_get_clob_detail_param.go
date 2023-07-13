package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClobDetailParam 获取clob数据的请求参数。
type GetClobDetailParam struct {

	// 环境id。
	EnvId int64 `json:"env_id"`

	// clobId。
	ClobId string `json:"clob_id"`
}

func (o GetClobDetailParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClobDetailParam struct{}"
	}

	return strings.Join([]string{"GetClobDetailParam", string(data)}, " ")
}
