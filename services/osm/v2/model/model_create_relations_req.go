package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRelationsReq struct {

	// 要关联的工单id列表，最多3个
	RelatedIdList []string `json:"related_id_list"`

	// 组id
	GroupId string `json:"group_id"`
}

func (o CreateRelationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRelationsReq struct{}"
	}

	return strings.Join([]string{"CreateRelationsReq", string(data)}, " ")
}
