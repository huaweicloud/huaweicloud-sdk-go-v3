package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMolDockingJobResponse Response Object
type CreateMolDockingJobResponse struct {

	// 对接打分结果。
	VinaScore *float32 `json:"vina_score,omitempty"`

	// 对接构象。
	Pose           *string `json:"pose,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMolDockingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMolDockingJobResponse struct{}"
	}

	return strings.Join([]string{"CreateMolDockingJobResponse", string(data)}, " ")
}
