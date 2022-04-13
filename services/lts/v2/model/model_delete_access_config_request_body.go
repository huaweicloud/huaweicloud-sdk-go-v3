package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除日志接入请求体
type DeleteAccessConfigRequestBody struct {
	// 日志接入ID列表

	AccessConfigIdList []string `json:"access_config_id_list"`
}

func (o DeleteAccessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteAccessConfigRequestBody", string(data)}, " ")
}
