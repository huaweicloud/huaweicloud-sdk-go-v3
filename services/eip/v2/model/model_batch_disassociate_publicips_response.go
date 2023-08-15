package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociatePublicipsResponse Response Object
type BatchDisassociatePublicipsResponse struct {

	// job_id列表，此job信息不保存在数据库中，不能同过组件查询到。
	JobIds         *[]string `json:"job_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDisassociatePublicipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociatePublicipsResponse struct{}"
	}

	return strings.Join([]string{"BatchDisassociatePublicipsResponse", string(data)}, " ")
}
