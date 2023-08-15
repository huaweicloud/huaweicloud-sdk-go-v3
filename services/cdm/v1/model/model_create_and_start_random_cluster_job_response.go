package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAndStartRandomClusterJobResponse Response Object
type CreateAndStartRandomClusterJobResponse struct {

	// 作业运行信息，请参见submission参数说明
	Submissions    *[]StartJobSubmission `json:"submissions,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CreateAndStartRandomClusterJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndStartRandomClusterJobResponse struct{}"
	}

	return strings.Join([]string{"CreateAndStartRandomClusterJobResponse", string(data)}, " ")
}
