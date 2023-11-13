package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDiagnosisTaskResponse Response Object
type DeleteDiagnosisTaskResponse struct {

	// 删除实例诊断结果
	Results        *[]Item `json:"results,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDiagnosisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDiagnosisTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteDiagnosisTaskResponse", string(data)}, " ")
}
