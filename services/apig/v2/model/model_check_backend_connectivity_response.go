package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CheckBackendConnectivityResponse struct {
	// 后端服务连通性检测结果  SUCCESS - 连通  FAILED - 不连通

	CheckResult *CheckBackendConnectivityResponseCheckResult `json:"check_result,omitempty"`
	// 后端服务连通检测失败的列表

	Failures       *[]string `json:"failures,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CheckBackendConnectivityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckBackendConnectivityResponse struct{}"
	}

	return strings.Join([]string{"CheckBackendConnectivityResponse", string(data)}, " ")
}

type CheckBackendConnectivityResponseCheckResult struct {
	value string
}

type CheckBackendConnectivityResponseCheckResultEnum struct {
	SUCCESS CheckBackendConnectivityResponseCheckResult
	FAILED  CheckBackendConnectivityResponseCheckResult
}

func GetCheckBackendConnectivityResponseCheckResultEnum() CheckBackendConnectivityResponseCheckResultEnum {
	return CheckBackendConnectivityResponseCheckResultEnum{
		SUCCESS: CheckBackendConnectivityResponseCheckResult{
			value: "SUCCESS",
		},
		FAILED: CheckBackendConnectivityResponseCheckResult{
			value: "FAILED",
		},
	}
}

func (c CheckBackendConnectivityResponseCheckResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckBackendConnectivityResponseCheckResult) UnmarshalJSON(b []byte) error {
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
