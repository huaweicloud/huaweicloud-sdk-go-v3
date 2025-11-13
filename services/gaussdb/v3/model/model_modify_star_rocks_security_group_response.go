package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyStarRocksSecurityGroupResponse Response Object
type ModifyStarRocksSecurityGroupResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyStarRocksSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyStarRocksSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"ModifyStarRocksSecurityGroupResponse", string(data)}, " ")
}
