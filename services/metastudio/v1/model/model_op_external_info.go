package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpExternalInfo 操作日志附加信息
type OpExternalInfo struct {

	// 审核详情id列表
	ReviewIdList *[]string `json:"review_id_list,omitempty"`
}

func (o OpExternalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExternalInfo struct{}"
	}

	return strings.Join([]string{"OpExternalInfo", string(data)}, " ")
}
