package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateResponse Response Object
type CreateTemplateResponse struct {

	// 任务运行结果
	Status *string `json:"status,omitempty"`

	Result         *StatusSuccessResultWithUuidResult `json:"result,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o CreateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplateResponse", string(data)}, " ")
}
