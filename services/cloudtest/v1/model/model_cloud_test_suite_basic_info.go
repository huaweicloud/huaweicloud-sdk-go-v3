package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudTestSuiteBasicInfo struct {
	CaseOperationInfo *CloudTestCaseOperationInfo `json:"caseOperationInfo,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 创建人ID
	CreateUserId *string `json:"create_user_id,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 数据类型：0为测试套，1为文件夹，cloudTest前台传入
	DocType *int32 `json:"doc_type,omitempty"`

	// 测试套状态
	ExecuteStatus *int32 `json:"execute_status,omitempty"`

	// 执行总次数
	ExecuteTimes *int32 `json:"execute_times,omitempty"`

	// 执行类型：0为冒烟测试，1为定时执行
	ExecuteType *int32 `json:"execute_type,omitempty"`

	// 执行方式：1为串行，2为并行，与echo的executeModel字段相同
	ExecuteWay *string `json:"execute_way,omitempty"`

	// 测试套超期状态
	ExpirationStatus *int32 `json:"expiration_status,omitempty"`

	// 参数配置
	ExtParam *string `json:"extParam,omitempty"`

	// 唯一ID，主键
	Id *string `json:"id,omitempty"`

	// 测试计划Uri，TMSS需要此值
	IteratorVersionUri *string `json:"iterator_version_uri,omitempty"`

	// 模块Id
	ModuleId *string `json:"moduleId,omitempty"`

	// 模块名称
	ModuleName *string `json:"moduleName,omitempty"`

	// 测试套名称，与echo的name字段相同
	Name *string `json:"name,omitempty"`

	// 目录Id
	NodeId *string `json:"nodeId,omitempty"`

	// 处理者ID
	OwnerId *string `json:"ownerId,omitempty"`

	// 测试计划id，可为空
	PlanId *string `json:"planId,omitempty"`

	// 计划结束时间
	PlanEndTimestamp *int64 `json:"plan_end_timestamp,omitempty"`

	// 计划开始时间
	PlanStartTimestamp *int64 `json:"plan_start_timestamp,omitempty"`

	// 项目ID
	ProjectId *int64 `json:"projectId,omitempty"`

	// 项目UUID，与echo的testServiceId字段相同
	ProjectUUId *string `json:"projectUUId,omitempty"`

	// 版本号
	ReleaseDev *string `json:"releaseDev,omitempty"`

	// 测试套执行结果
	Result *int32 `json:"result,omitempty"`

	// 测试套状态
	Status *int32 `json:"status,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	// 测试套id，更新时需要同时传id、testSuiteId，2个字段值相同，与echo的taskId字段相同
	TestSuiteId *string `json:"testSuiteId,omitempty"`

	// 编号
	TestSuiteNumber *string `json:"testSuiteNumber,omitempty"`

	// 测试套类型：0为功能测试，1为接口测试，6为Pistar，cloudTest前台传入
	Type *int32 `json:"type,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`

	// 更新人ID
	UpdateUserId *string `json:"update_user_id,omitempty"`
}

func (o CloudTestSuiteBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudTestSuiteBasicInfo struct{}"
	}

	return strings.Join([]string{"CloudTestSuiteBasicInfo", string(data)}, " ")
}
