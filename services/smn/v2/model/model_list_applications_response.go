package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApplicationsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 返回的Application个数。该参数不受offset和limit影响，即返回的是您账户下所有的Application个数。
	ApplicationCount *int32 `json:"application_count,omitempty" xml:"application_count"`

	Applications   *[]ApplicationItem `json:"applications,omitempty" xml:"applications"`
	HttpStatusCode int                `json:"-"`
}

func (o ListApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationsResponse", string(data)}, " ")
}
