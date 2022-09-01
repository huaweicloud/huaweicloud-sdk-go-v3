package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAreaDetailResponse struct {

	// 时间戳及相应时间的指标数值
	DataList *[]AreaDetail `json:"data_list,omitempty" xml:"data_list"`

	XRequestId     *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAreaDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAreaDetailResponse struct{}"
	}

	return strings.Join([]string{"ListAreaDetailResponse", string(data)}, " ")
}
