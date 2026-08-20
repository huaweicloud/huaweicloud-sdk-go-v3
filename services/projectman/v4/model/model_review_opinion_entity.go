package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReviewOpinionEntity struct {

	// 评审意见对象类型，固定为Opinion。
	Category *string `json:"category,omitempty"`

	// 评审意见对象关联的变更对象ID。
	CoId *string `json:"co_id,omitempty"`

	CreatedBy *UserEntity `json:"created_by,omitempty"`

	// 评审意见创建时间。
	CreatedDate *string `json:"created_date,omitempty"`

	CurrOwner *UserEntity `json:"curr_owner,omitempty"`

	// 评审意见对象ID。
	Id *string `json:"id,omitempty"`

	// 评审意见最后修改时间。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 评审意见。
	ReviewComments *string `json:"review_comments,omitempty"`
}

func (o ReviewOpinionEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewOpinionEntity struct{}"
	}

	return strings.Join([]string{"ReviewOpinionEntity", string(data)}, " ")
}
