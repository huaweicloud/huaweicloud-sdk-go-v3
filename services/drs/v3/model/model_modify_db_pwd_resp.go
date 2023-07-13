package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDbPwdResp 修改数据库密码返回体
type ModifyDbPwdResp struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 源库：so，目标库：ta
	EndPointType *string `json:"end_point_type,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ModifyDbPwdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbPwdResp struct{}"
	}

	return strings.Join([]string{"ModifyDbPwdResp", string(data)}, " ")
}
