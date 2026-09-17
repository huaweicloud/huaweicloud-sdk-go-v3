package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareSlowLogTemplatesResponse Response Object
type CompareSlowLogTemplatesResponse struct {

	// 模板数据对比结果列表
	Contrasts      *[]SlowLogTplContrast `json:"contrasts,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CompareSlowLogTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareSlowLogTemplatesResponse struct{}"
	}

	return strings.Join([]string{"CompareSlowLogTemplatesResponse", string(data)}, " ")
}
