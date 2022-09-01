package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeGaussMySqlProxySpecificationResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeGaussMySqlProxySpecificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeGaussMySqlProxySpecificationResponse struct{}"
	}

	return strings.Join([]string{"ChangeGaussMySqlProxySpecificationResponse", string(data)}, " ")
}
