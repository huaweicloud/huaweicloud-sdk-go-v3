package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRelationsReq struct {

	// 要关联的工单id列表，最多3个
	RelatedIdList []string `json:"related_id_list"`

	// 华为云IAM组id，操作查询同组其他工单时，该id必传
	GroupId *string `json:"group_id,omitempty"`
}

func (o CreateRelationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRelationsReq struct{}"
	}

	return strings.Join([]string{"CreateRelationsReq", string(data)}, " ")
}
