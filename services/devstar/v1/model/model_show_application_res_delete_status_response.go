package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowApplicationResDeleteStatusResponse struct {

	// 代码仓删除状态,deleted:删除成功,failed:删除失败,going:正在删除中
	RepoStatus *ShowApplicationResDeleteStatusResponseRepoStatus `json:"repo_status,omitempty"`

	// 流水线删除状态
	PipelineStatus *[]PipelineDeleteStatus `json:"pipeline_status,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowApplicationResDeleteStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationResDeleteStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowApplicationResDeleteStatusResponse", string(data)}, " ")
}

type ShowApplicationResDeleteStatusResponseRepoStatus struct {
	value string
}

type ShowApplicationResDeleteStatusResponseRepoStatusEnum struct {
	DELETED ShowApplicationResDeleteStatusResponseRepoStatus
	FAILED  ShowApplicationResDeleteStatusResponseRepoStatus
	GOING   ShowApplicationResDeleteStatusResponseRepoStatus
}

func GetShowApplicationResDeleteStatusResponseRepoStatusEnum() ShowApplicationResDeleteStatusResponseRepoStatusEnum {
	return ShowApplicationResDeleteStatusResponseRepoStatusEnum{
		DELETED: ShowApplicationResDeleteStatusResponseRepoStatus{
			value: "deleted",
		},
		FAILED: ShowApplicationResDeleteStatusResponseRepoStatus{
			value: "failed",
		},
		GOING: ShowApplicationResDeleteStatusResponseRepoStatus{
			value: "going",
		},
	}
}

func (c ShowApplicationResDeleteStatusResponseRepoStatus) Value() string {
	return c.value
}

func (c ShowApplicationResDeleteStatusResponseRepoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplicationResDeleteStatusResponseRepoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
