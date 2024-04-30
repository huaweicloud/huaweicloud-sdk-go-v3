package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OperationPreferences 资源栈集操作（stack_set_operation）的部署策略。该参数只在指定的单次操作中生效。  当用户不指定该参数时，默认的操作部署策略为区域（region）内资源栈实例串行部署，即每次只执行一个资源栈实例，区域（region）间随机且串行部署，执行完一个region下的全部资源栈实例后，才会选择另一个region部署，容错次数默认为0。  该参数可以在生成资源栈集操作的四个API中指定：  创建资源栈实例（CreateStackInstance），部署资源栈集（DeployStackSet），更新资源栈实例（UpdateStackInstance），删除资源栈实例（DeleteStackInstance）
type OperationPreferences struct {

	// 部署资源栈实例时区域（region）的执行策略，分为两种，SEQUENTIAL和PARALLEL，区分大小写，默认值为SEQUENTIAL  详细介绍：  * `SEQUENTIAL`：顺序执行，执行完一个region下的全部资源栈实例后再去执行另一个region。默认顺序执行。  * `PARALLEL`：并发执行，并发部署所有指定区域的资源栈实例。
	RegionConcurrencyType *OperationPreferencesRegionConcurrencyType `json:"region_concurrency_type,omitempty"`

	// 区域（region）部署顺序。只有当用户指定region_concurrency_type为SEQUENTIAL时才会允许指定该参数。用户指定部署region的顺序，不允许出现资源栈集管理之外的region。  如果不指定，实际部署region顺序随机。部署顺序仅在当次部署时生效，应该包含且仅包含本次部署的所有region。
	RegionOrder *[]string `json:"region_order,omitempty"`

	// 容错次数。用户定义在每个区域（region）下，允许部署失败的资源栈实例数量。该参数取值默认为0，限定0和正整数。  如果定义region顺序执行（region_concurrency_type值为SEQUENTIAL），在某个region超过容错次数时，资源栈集会取消所有状态仍处于WAIT_IN_PROGRESS状态的实例。被取消的实例状态最终变为CANCEL_COMPLETE；  如果是region并行执行（region_concurrency_type值为PARALLEL），在某个region超过容错次数时，资源栈集只会取消该region下所有处于WAIT_IN_PROGRESS状态的实例。被取消的实例状态最终变为CANCEL_COMPLETE。  对处于OPERATION_IN_PROGRESS，或已经部署完成，即处于OPERATION_COMPLETE或者OPERATION_FAILED状态的资源栈实例，不受影响，状态不变。  failure_tolerance_count  和 failure_tolerance_percentage 仅能有一个存在。
	FailureToleranceCount *int64 `json:"failure_tolerance_count,omitempty"`

	// 容错百分比。定义每个区域（region）下，允许部署失败的资源栈实例数占该region下所有资源栈实例数的百分比。该参数取值默认为0，限定0和正整数。  通过容错百分比*资源栈实例数，并向下取整，得到实际容错次数。  failure_tolerance_count  和 failure_tolerance_percentage 仅能有一个存在。
	FailureTolerancePercentage *int64 `json:"failure_tolerance_percentage,omitempty"`

	// 每个区域（region）下可同时部署资源栈实例的最大账户数。该参数取值默认为1，限定为正整数。  最大并发账户数最多比容错次数多1。如果用户指定failure_tolerance_percentage，最大并发账户数最多比 failure_tolerance_percentage * 资源栈实例数多1。保证部署在所需的容错级别停止。  max_concurrent_count 和 max_concurrent_percentage 仅能有一个存在。
	MaxConcurrentCount *int64 `json:"max_concurrent_count,omitempty"`

	// 最大并发账户百分比，每个区域（region）中可同时部署的资源栈实例的最大账户百分比。该参数取值默认为1，限定正整数。  RFS根据百分比 *（每个region下资源栈实例数）得到的值，再向下取整，得到实际最大并发账户数。如果实际最大并发账户数向下取整值为0时，则默认选择最大并发账户数为1。  通过百分比计算得到的实际最大并发账户数最多比容错次数多1。如果用户指定failure_tolerance_percentage，实际最大并发账户数最多比 failure_tolerance_percentage * 资源栈实例数多1。保证部署在所需的容错级别停止。  max_concurrent_count 和 max_concurrent_percentage 仅能有一个存在。
	MaxConcurrentPercentage *int64 `json:"max_concurrent_percentage,omitempty"`

	// 资源栈集操作部署的失败容忍模式，分为两种，STRICT_FAILURE_TOLERANCE和SOFT_FAILURE_TOLERANCE，区分大小写，默认值为STRICT_FAILURE_TOLERANCE。  详细介绍：  * `STRICT_FAILURE_TOLERANCE`：此选项会动态降低并发级别，以确保同region下部署失败的账户数量永远不超过 failure_tolerance_count + 1。当用户指定failure_tolerance_percentage时，确保同region下部署失败的账户数量不超过 failure_tolerance_percentage * 资源栈实例数 + 1。  * 初始实际最大并发数为max_concurrent_count，如果用户指定的是max_concurrent_percentage，则初始实际最大并发数为 max_concurrent_percentage * 资源栈实例数，随后，实际最大并发数会根据失败次数增加而减少。  * `SOFT_FAILURE_TOLERANCE`：此选项将failure_tolerance_count (failure_tolerance_percentage) 与实际并发数分离开。该参数允许资源栈集操作始终以指定的 max_concurrent_count 或 max_concurrent_percentage 操作资源栈实例。  * 此时不保证资源栈实例失败总数小于 failure_tolerance_count + 1，如果用户指定的是failure_tolerance_percentage的值，则不保证资源栈实例失败总数小于 failure_tolerance_percentage * 资源栈实例数 + 1。
	FailureToleranceMode *OperationPreferencesFailureToleranceMode `json:"failure_tolerance_mode,omitempty"`
}

func (o OperationPreferences) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationPreferences struct{}"
	}

	return strings.Join([]string{"OperationPreferences", string(data)}, " ")
}

type OperationPreferencesRegionConcurrencyType struct {
	value string
}

type OperationPreferencesRegionConcurrencyTypeEnum struct {
	SEQUENTIAL OperationPreferencesRegionConcurrencyType
	PARALLEL   OperationPreferencesRegionConcurrencyType
}

func GetOperationPreferencesRegionConcurrencyTypeEnum() OperationPreferencesRegionConcurrencyTypeEnum {
	return OperationPreferencesRegionConcurrencyTypeEnum{
		SEQUENTIAL: OperationPreferencesRegionConcurrencyType{
			value: "SEQUENTIAL",
		},
		PARALLEL: OperationPreferencesRegionConcurrencyType{
			value: "PARALLEL",
		},
	}
}

func (c OperationPreferencesRegionConcurrencyType) Value() string {
	return c.value
}

func (c OperationPreferencesRegionConcurrencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationPreferencesRegionConcurrencyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type OperationPreferencesFailureToleranceMode struct {
	value string
}

type OperationPreferencesFailureToleranceModeEnum struct {
	STRICT_FAILURE_TOLERANCE OperationPreferencesFailureToleranceMode
	SOFT_FAILURE_TOLERANCE   OperationPreferencesFailureToleranceMode
}

func GetOperationPreferencesFailureToleranceModeEnum() OperationPreferencesFailureToleranceModeEnum {
	return OperationPreferencesFailureToleranceModeEnum{
		STRICT_FAILURE_TOLERANCE: OperationPreferencesFailureToleranceMode{
			value: "STRICT_FAILURE_TOLERANCE",
		},
		SOFT_FAILURE_TOLERANCE: OperationPreferencesFailureToleranceMode{
			value: "SOFT_FAILURE_TOLERANCE",
		},
	}
}

func (c OperationPreferencesFailureToleranceMode) Value() string {
	return c.value
}

func (c OperationPreferencesFailureToleranceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationPreferencesFailureToleranceMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
