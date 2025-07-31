package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBaselineWhiteListRequestBody 删除基线白名单请求体信息
type DeleteBaselineWhiteListRequestBody struct {

	// 基线白名单ID列表
	IdList []string `json:"id_list"`
}

func (o DeleteBaselineWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBaselineWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteBaselineWhiteListRequestBody", string(data)}, " ")
}
