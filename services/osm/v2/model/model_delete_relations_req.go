package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteRelationsReq struct {
	// 要解除关联的工单id

	RelatedId string `json:"related_id"`
	// 组id

	GroupId string `json:"group_id"`
}

func (o DeleteRelationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRelationsReq struct{}"
	}

	return strings.Join([]string{"DeleteRelationsReq", string(data)}, " ")
}
