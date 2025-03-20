package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutCaseExtParamReq struct {

	// 华为云IAM组id，操作查询同组其他工单时，该id必传
	GroupId *string `json:"group_id,omitempty"`

	// 消息id
	MessageId *string `json:"message_id,omitempty"`

	// 扩展参数
	ExtendsMap map[string]interface{} `json:"extends_map,omitempty"`
}

func (o PutCaseExtParamReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCaseExtParamReq struct{}"
	}

	return strings.Join([]string{"PutCaseExtParamReq", string(data)}, " ")
}
