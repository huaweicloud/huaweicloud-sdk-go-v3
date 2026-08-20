package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutIpdChangeReviewFormV2Request Request Object
type PutIpdChangeReviewFormV2Request struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 评审单ID，评审单唯一标识。通过查询评审单列表（BR/GR）接口获取，响应消息体中的id字段的值就是评审单ID。
	ReviewId string `json:"review_id"`

	Body *ReviewUpdateBodyV2 `json:"body,omitempty"`
}

func (o PutIpdChangeReviewFormV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutIpdChangeReviewFormV2Request struct{}"
	}

	return strings.Join([]string{"PutIpdChangeReviewFormV2Request", string(data)}, " ")
}
