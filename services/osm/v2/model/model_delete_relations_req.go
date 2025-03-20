package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteRelationsReq struct {

	// 要解除关联的工单id
	RelatedId string `json:"related_id"`

	// 华为云IAM组id，操作查询同组其他工单时，该id必传
	GroupId *string `json:"group_id,omitempty"`
}

func (o DeleteRelationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRelationsReq struct{}"
	}

	return strings.Join([]string{"DeleteRelationsReq", string(data)}, " ")
}
