package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteDomainSetsDto struct {

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 域名组id
	SetIds *[]string `json:"set_ids,omitempty"`
}

func (o BatchDeleteDomainSetsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDomainSetsDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteDomainSetsDto", string(data)}, " ")
}
