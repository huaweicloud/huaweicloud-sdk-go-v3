package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlInstanceDetailInfoResponse Response Object
type ListGaussMySqlInstanceDetailInfoResponse struct {

	// 实例详情。
	Instances      *[]MysqlInstanceInfoDetail `json:"instances,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListGaussMySqlInstanceDetailInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlInstanceDetailInfoResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlInstanceDetailInfoResponse", string(data)}, " ")
}
