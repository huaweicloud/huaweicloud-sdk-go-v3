package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResInstanceInfoResponse Response Object
type ShowResInstanceInfoResponse struct {

	// 企业项目列表
	Resources *[]ListEnterpriseResourceResult `json:"resources,omitempty"`

	// 资源数
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowResInstanceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResInstanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowResInstanceInfoResponse", string(data)}, " ")
}
