package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rgc/v1/model"
)

type RgcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewRgcClient(hcClient *httpclient.HcHttpClient) *RgcClient {
	return &RgcClient{HcClient: hcClient}
}

func RgcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateBestPracticeDetect 检测八大场景治理成熟度
//
// 检测八大场景治理成熟度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CreateBestPracticeDetect(request *model.CreateBestPracticeDetectRequest) (*model.CreateBestPracticeDetectResponse, error) {
	requestDef := GenReqDefForCreateBestPracticeDetect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBestPracticeDetectResponse), nil
	}
}

// CreateBestPracticeDetectInvoker 检测八大场景治理成熟度
func (c *RgcClient) CreateBestPracticeDetectInvoker(request *model.CreateBestPracticeDetectRequest) *CreateBestPracticeDetectInvoker {
	requestDef := GenReqDefForCreateBestPracticeDetect()
	return &CreateBestPracticeDetectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBestPracticeAccountInfo 查询治理成熟度的账号详情
//
// 查询治理成熟度的账号详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowBestPracticeAccountInfo(request *model.ShowBestPracticeAccountInfoRequest) (*model.ShowBestPracticeAccountInfoResponse, error) {
	requestDef := GenReqDefForShowBestPracticeAccountInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBestPracticeAccountInfoResponse), nil
	}
}

// ShowBestPracticeAccountInfoInvoker 查询治理成熟度的账号详情
func (c *RgcClient) ShowBestPracticeAccountInfoInvoker(request *model.ShowBestPracticeAccountInfoRequest) *ShowBestPracticeAccountInfoInvoker {
	requestDef := GenReqDefForShowBestPracticeAccountInfo()
	return &ShowBestPracticeAccountInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBestPracticeDetails 查询最近一次成功的治理成熟度检测的详情
//
// 查询最近一次成功的治理成熟度检测的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowBestPracticeDetails(request *model.ShowBestPracticeDetailsRequest) (*model.ShowBestPracticeDetailsResponse, error) {
	requestDef := GenReqDefForShowBestPracticeDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBestPracticeDetailsResponse), nil
	}
}

// ShowBestPracticeDetailsInvoker 查询最近一次成功的治理成熟度检测的详情
func (c *RgcClient) ShowBestPracticeDetailsInvoker(request *model.ShowBestPracticeDetailsRequest) *ShowBestPracticeDetailsInvoker {
	requestDef := GenReqDefForShowBestPracticeDetails()
	return &ShowBestPracticeDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBestPracticeOverview 查询最近一次成功的治理成熟度检测的总览
//
// 查询最近一次成功的治理成熟度检测的总览。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowBestPracticeOverview(request *model.ShowBestPracticeOverviewRequest) (*model.ShowBestPracticeOverviewResponse, error) {
	requestDef := GenReqDefForShowBestPracticeOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBestPracticeOverviewResponse), nil
	}
}

// ShowBestPracticeOverviewInvoker 查询最近一次成功的治理成熟度检测的总览
func (c *RgcClient) ShowBestPracticeOverviewInvoker(request *model.ShowBestPracticeOverviewRequest) *ShowBestPracticeOverviewInvoker {
	requestDef := GenReqDefForShowBestPracticeOverview()
	return &ShowBestPracticeOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBestPracticeStatus 查询最近一次的治理成熟度检测的状态
//
// 查询最近一次的治理成熟度检测的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowBestPracticeStatus(request *model.ShowBestPracticeStatusRequest) (*model.ShowBestPracticeStatusResponse, error) {
	requestDef := GenReqDefForShowBestPracticeStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBestPracticeStatusResponse), nil
	}
}

// ShowBestPracticeStatusInvoker 查询最近一次的治理成熟度检测的状态
func (c *RgcClient) ShowBestPracticeStatusInvoker(request *model.ShowBestPracticeStatusRequest) *ShowBestPracticeStatusInvoker {
	requestDef := GenReqDefForShowBestPracticeStatus()
	return &ShowBestPracticeStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableControl 关闭控制策略
//
// 关闭组织下的某个单元的某个控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DisableControl(request *model.DisableControlRequest) (*model.DisableControlResponse, error) {
	requestDef := GenReqDefForDisableControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableControlResponse), nil
	}
}

// DisableControlInvoker 关闭控制策略
func (c *RgcClient) DisableControlInvoker(request *model.DisableControlRequest) *DisableControlInvoker {
	requestDef := GenReqDefForDisableControl()
	return &DisableControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableControl 开启控制策略
//
// 给组织下的某个单元开启某个控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) EnableControl(request *model.EnableControlRequest) (*model.EnableControlResponse, error) {
	requestDef := GenReqDefForEnableControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableControlResponse), nil
	}
}

// EnableControlInvoker 开启控制策略
func (c *RgcClient) EnableControlInvoker(request *model.EnableControlRequest) *EnableControlInvoker {
	requestDef := GenReqDefForEnableControl()
	return &EnableControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigRuleCompliances 查询纳管账号的Config规则合规性信息
//
// 查询纳管账号的Config规则合规性信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListConfigRuleCompliances(request *model.ListConfigRuleCompliancesRequest) (*model.ListConfigRuleCompliancesResponse, error) {
	requestDef := GenReqDefForListConfigRuleCompliances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigRuleCompliancesResponse), nil
	}
}

// ListConfigRuleCompliancesInvoker 查询纳管账号的Config规则合规性信息
func (c *RgcClient) ListConfigRuleCompliancesInvoker(request *model.ListConfigRuleCompliancesRequest) *ListConfigRuleCompliancesInvoker {
	requestDef := GenReqDefForListConfigRuleCompliances()
	return &ListConfigRuleCompliancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControlViolations 列出不合规信息
//
// 列出组织里所有不合规的资源信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlViolations(request *model.ListControlViolationsRequest) (*model.ListControlViolationsResponse, error) {
	requestDef := GenReqDefForListControlViolations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlViolationsResponse), nil
	}
}

// ListControlViolationsInvoker 列出不合规信息
func (c *RgcClient) ListControlViolationsInvoker(request *model.ListControlViolationsRequest) *ListControlViolationsInvoker {
	requestDef := GenReqDefForListControlViolations()
	return &ListControlViolationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControls 列出控制策略
//
// 列出RGC服务里所有的预置控制策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControls(request *model.ListControlsRequest) (*model.ListControlsResponse, error) {
	requestDef := GenReqDefForListControls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsResponse), nil
	}
}

// ListControlsInvoker 列出控制策略
func (c *RgcClient) ListControlsInvoker(request *model.ListControlsRequest) *ListControlsInvoker {
	requestDef := GenReqDefForListControls()
	return &ListControlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControlsForAccount 列出纳管账号下开启的控制策略
//
// 列出组织里某个纳管账号开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlsForAccount(request *model.ListControlsForAccountRequest) (*model.ListControlsForAccountResponse, error) {
	requestDef := GenReqDefForListControlsForAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsForAccountResponse), nil
	}
}

// ListControlsForAccountInvoker 列出纳管账号下开启的控制策略
func (c *RgcClient) ListControlsForAccountInvoker(request *model.ListControlsForAccountRequest) *ListControlsForAccountInvoker {
	requestDef := GenReqDefForListControlsForAccount()
	return &ListControlsForAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListControlsForOrganizationalUnit 列出注册OU下开启的控制策略
//
// 列出组织里某个注册OU开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListControlsForOrganizationalUnit(request *model.ListControlsForOrganizationalUnitRequest) (*model.ListControlsForOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForListControlsForOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListControlsForOrganizationalUnitResponse), nil
	}
}

// ListControlsForOrganizationalUnitInvoker 列出注册OU下开启的控制策略
func (c *RgcClient) ListControlsForOrganizationalUnitInvoker(request *model.ListControlsForOrganizationalUnitRequest) *ListControlsForOrganizationalUnitInvoker {
	requestDef := GenReqDefForListControlsForOrganizationalUnit()
	return &ListControlsForOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDriftDetails 列出漂移信息
//
// 列出Landing Zone的所有漂移详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListDriftDetails(request *model.ListDriftDetailsRequest) (*model.ListDriftDetailsResponse, error) {
	requestDef := GenReqDefForListDriftDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDriftDetailsResponse), nil
	}
}

// ListDriftDetailsInvoker 列出漂移信息
func (c *RgcClient) ListDriftDetailsInvoker(request *model.ListDriftDetailsRequest) *ListDriftDetailsInvoker {
	requestDef := GenReqDefForListDriftDetails()
	return &ListDriftDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnabledControls 列出开启的控制策略
//
// 列出组织里开启的所有控制策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListEnabledControls(request *model.ListEnabledControlsRequest) (*model.ListEnabledControlsResponse, error) {
	requestDef := GenReqDefForListEnabledControls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnabledControlsResponse), nil
	}
}

// ListEnabledControlsInvoker 列出开启的控制策略
func (c *RgcClient) ListEnabledControlsInvoker(request *model.ListEnabledControlsRequest) *ListEnabledControlsInvoker {
	requestDef := GenReqDefForListEnabledControls()
	return &ListEnabledControlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExternalConfigRuleCompliances 查询纳管账号的外部Config规则合规性信息
//
// 查询纳管账号的外部Config规则合规性信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListExternalConfigRuleCompliances(request *model.ListExternalConfigRuleCompliancesRequest) (*model.ListExternalConfigRuleCompliancesResponse, error) {
	requestDef := GenReqDefForListExternalConfigRuleCompliances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExternalConfigRuleCompliancesResponse), nil
	}
}

// ListExternalConfigRuleCompliancesInvoker 查询纳管账号的外部Config规则合规性信息
func (c *RgcClient) ListExternalConfigRuleCompliancesInvoker(request *model.ListExternalConfigRuleCompliancesRequest) *ListExternalConfigRuleCompliancesInvoker {
	requestDef := GenReqDefForListExternalConfigRuleCompliances()
	return &ListExternalConfigRuleCompliancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComplianceStatusForAccount 查询纳管账号的合规状态
//
// 查询组织里某个纳管账号的资源合规状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowComplianceStatusForAccount(request *model.ShowComplianceStatusForAccountRequest) (*model.ShowComplianceStatusForAccountResponse, error) {
	requestDef := GenReqDefForShowComplianceStatusForAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComplianceStatusForAccountResponse), nil
	}
}

// ShowComplianceStatusForAccountInvoker 查询纳管账号的合规状态
func (c *RgcClient) ShowComplianceStatusForAccountInvoker(request *model.ShowComplianceStatusForAccountRequest) *ShowComplianceStatusForAccountInvoker {
	requestDef := GenReqDefForShowComplianceStatusForAccount()
	return &ShowComplianceStatusForAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComplianceStatusForOrganizationalUnit 查询注册OU的合规状态
//
// 查询组织里某个注册OU下所有纳管账号的资源合规状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowComplianceStatusForOrganizationalUnit(request *model.ShowComplianceStatusForOrganizationalUnitRequest) (*model.ShowComplianceStatusForOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForShowComplianceStatusForOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComplianceStatusForOrganizationalUnitResponse), nil
	}
}

// ShowComplianceStatusForOrganizationalUnitInvoker 查询注册OU的合规状态
func (c *RgcClient) ShowComplianceStatusForOrganizationalUnitInvoker(request *model.ShowComplianceStatusForOrganizationalUnitRequest) *ShowComplianceStatusForOrganizationalUnitInvoker {
	requestDef := GenReqDefForShowComplianceStatusForOrganizationalUnit()
	return &ShowComplianceStatusForOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControl 查询控制策略详细信息
//
// 查询单个预置的控制策略详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControl(request *model.ShowControlRequest) (*model.ShowControlResponse, error) {
	requestDef := GenReqDefForShowControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlResponse), nil
	}
}

// ShowControlInvoker 查询控制策略详细信息
func (c *RgcClient) ShowControlInvoker(request *model.ShowControlRequest) *ShowControlInvoker {
	requestDef := GenReqDefForShowControl()
	return &ShowControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControlOperate 查询控制策略操作状态
//
// 根据操作ID查询返回指定ID的操作状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControlOperate(request *model.ShowControlOperateRequest) (*model.ShowControlOperateResponse, error) {
	requestDef := GenReqDefForShowControlOperate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlOperateResponse), nil
	}
}

// ShowControlOperateInvoker 查询控制策略操作状态
func (c *RgcClient) ShowControlOperateInvoker(request *model.ShowControlOperateRequest) *ShowControlOperateInvoker {
	requestDef := GenReqDefForShowControlOperate()
	return &ShowControlOperateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControlsForAccount 查询纳管账号开启的控制策略
//
// 查询组织里某个纳管账号下开启的某个控制策略的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControlsForAccount(request *model.ShowControlsForAccountRequest) (*model.ShowControlsForAccountResponse, error) {
	requestDef := GenReqDefForShowControlsForAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlsForAccountResponse), nil
	}
}

// ShowControlsForAccountInvoker 查询纳管账号开启的控制策略
func (c *RgcClient) ShowControlsForAccountInvoker(request *model.ShowControlsForAccountRequest) *ShowControlsForAccountInvoker {
	requestDef := GenReqDefForShowControlsForAccount()
	return &ShowControlsForAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowControlsForOrganizationalUnit 查询注册OU开启的控制策略
//
// 查询组织里某个注册OU下开启的某个控制策略的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowControlsForOrganizationalUnit(request *model.ShowControlsForOrganizationalUnitRequest) (*model.ShowControlsForOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForShowControlsForOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowControlsForOrganizationalUnitResponse), nil
	}
}

// ShowControlsForOrganizationalUnitInvoker 查询注册OU开启的控制策略
func (c *RgcClient) ShowControlsForOrganizationalUnitInvoker(request *model.ShowControlsForOrganizationalUnitRequest) *ShowControlsForOrganizationalUnitInvoker {
	requestDef := GenReqDefForShowControlsForOrganizationalUnit()
	return &ShowControlsForOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckLaunch 设置Landing Zone前检查
//
// 在设置Landing Zone之前，检查当前区域是否可以设置Landing Zone。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CheckLaunch(request *model.CheckLaunchRequest) (*model.CheckLaunchResponse, error) {
	requestDef := GenReqDefForCheckLaunch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckLaunchResponse), nil
	}
}

// CheckLaunchInvoker 设置Landing Zone前检查
func (c *RgcClient) CheckLaunchInvoker(request *model.CheckLaunchRequest) *CheckLaunchInvoker {
	requestDef := GenReqDefForCheckLaunch()
	return &CheckLaunchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLandingZone 删除Landing Zone
//
// 删除Landing Zone。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DeleteLandingZone(request *model.DeleteLandingZoneRequest) (*model.DeleteLandingZoneResponse, error) {
	requestDef := GenReqDefForDeleteLandingZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLandingZoneResponse), nil
	}
}

// DeleteLandingZoneInvoker 删除Landing Zone
func (c *RgcClient) DeleteLandingZoneInvoker(request *model.DeleteLandingZoneRequest) *DeleteLandingZoneInvoker {
	requestDef := GenReqDefForDeleteLandingZone()
	return &DeleteLandingZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetupLandingZone 设置Landing Zone
//
// 在当前区域创建或者更新Landing Zone。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) SetupLandingZone(request *model.SetupLandingZoneRequest) (*model.SetupLandingZoneResponse, error) {
	requestDef := GenReqDefForSetupLandingZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetupLandingZoneResponse), nil
	}
}

// SetupLandingZoneInvoker 设置Landing Zone
func (c *RgcClient) SetupLandingZoneInvoker(request *model.SetupLandingZoneRequest) *SetupLandingZoneInvoker {
	requestDef := GenReqDefForSetupLandingZone()
	return &SetupLandingZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvailableUpdates 查询Landing Zone可更新状态
//
// 查询Landing Zone当前是否需要升级更新。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowAvailableUpdates(request *model.ShowAvailableUpdatesRequest) (*model.ShowAvailableUpdatesResponse, error) {
	requestDef := GenReqDefForShowAvailableUpdates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvailableUpdatesResponse), nil
	}
}

// ShowAvailableUpdatesInvoker 查询Landing Zone可更新状态
func (c *RgcClient) ShowAvailableUpdatesInvoker(request *model.ShowAvailableUpdatesRequest) *ShowAvailableUpdatesInvoker {
	requestDef := GenReqDefForShowAvailableUpdates()
	return &ShowAvailableUpdatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHomeRegion 查询主区域
//
// 查询Landing Zone的主区域。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowHomeRegion(request *model.ShowHomeRegionRequest) (*model.ShowHomeRegionResponse, error) {
	requestDef := GenReqDefForShowHomeRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHomeRegionResponse), nil
	}
}

// ShowHomeRegionInvoker 查询主区域
func (c *RgcClient) ShowHomeRegionInvoker(request *model.ShowHomeRegionRequest) *ShowHomeRegionInvoker {
	requestDef := GenReqDefForShowHomeRegion()
	return &ShowHomeRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLandingZoneConfiguration 查询Landing Zone的配置
//
// 查询当前客户的Landing Zone的所有配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowLandingZoneConfiguration(request *model.ShowLandingZoneConfigurationRequest) (*model.ShowLandingZoneConfigurationResponse, error) {
	requestDef := GenReqDefForShowLandingZoneConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLandingZoneConfigurationResponse), nil
	}
}

// ShowLandingZoneConfigurationInvoker 查询Landing Zone的配置
func (c *RgcClient) ShowLandingZoneConfigurationInvoker(request *model.ShowLandingZoneConfigurationRequest) *ShowLandingZoneConfigurationInvoker {
	requestDef := GenReqDefForShowLandingZoneConfiguration()
	return &ShowLandingZoneConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLandingZoneIdentityCenter 查询当前客户的Identity Center用户信息
//
// 查询当前客户的Identity Center用户信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowLandingZoneIdentityCenter(request *model.ShowLandingZoneIdentityCenterRequest) (*model.ShowLandingZoneIdentityCenterResponse, error) {
	requestDef := GenReqDefForShowLandingZoneIdentityCenter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLandingZoneIdentityCenterResponse), nil
	}
}

// ShowLandingZoneIdentityCenterInvoker 查询当前客户的Identity Center用户信息
func (c *RgcClient) ShowLandingZoneIdentityCenterInvoker(request *model.ShowLandingZoneIdentityCenterRequest) *ShowLandingZoneIdentityCenterInvoker {
	requestDef := GenReqDefForShowLandingZoneIdentityCenter()
	return &ShowLandingZoneIdentityCenterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLandingZoneStatus 查询Landing Zone设置状态
//
// 查询Landing Zone的设置状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowLandingZoneStatus(request *model.ShowLandingZoneStatusRequest) (*model.ShowLandingZoneStatusResponse, error) {
	requestDef := GenReqDefForShowLandingZoneStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLandingZoneStatusResponse), nil
	}
}

// ShowLandingZoneStatusInvoker 查询Landing Zone设置状态
func (c *RgcClient) ShowLandingZoneStatusInvoker(request *model.ShowLandingZoneStatusRequest) *ShowLandingZoneStatusInvoker {
	requestDef := GenReqDefForShowLandingZoneStatus()
	return &ShowLandingZoneStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccount 创建账号
//
// 在组织里的某个注册OU下创建账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CreateAccount(request *model.CreateAccountRequest) (*model.CreateAccountResponse, error) {
	requestDef := GenReqDefForCreateAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccountResponse), nil
	}
}

// CreateAccountInvoker 创建账号
func (c *RgcClient) CreateAccountInvoker(request *model.CreateAccountRequest) *CreateAccountInvoker {
	requestDef := GenReqDefForCreateAccount()
	return &CreateAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateManagedOrganizationalUnit 创建OU
//
// 通过RGC服务在组织下创建OU，创建后的OU在RGC中状态为已注册。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CreateManagedOrganizationalUnit(request *model.CreateManagedOrganizationalUnitRequest) (*model.CreateManagedOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForCreateManagedOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManagedOrganizationalUnitResponse), nil
	}
}

// CreateManagedOrganizationalUnitInvoker 创建OU
func (c *RgcClient) CreateManagedOrganizationalUnitInvoker(request *model.CreateManagedOrganizationalUnitRequest) *CreateManagedOrganizationalUnitInvoker {
	requestDef := GenReqDefForCreateManagedOrganizationalUnit()
	return &CreateManagedOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteManagedOrganizationalUnits 删除注册OU
//
// 在组织里删除已注册OU。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DeleteManagedOrganizationalUnits(request *model.DeleteManagedOrganizationalUnitsRequest) (*model.DeleteManagedOrganizationalUnitsResponse, error) {
	requestDef := GenReqDefForDeleteManagedOrganizationalUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteManagedOrganizationalUnitsResponse), nil
	}
}

// DeleteManagedOrganizationalUnitsInvoker 删除注册OU
func (c *RgcClient) DeleteManagedOrganizationalUnitsInvoker(request *model.DeleteManagedOrganizationalUnitsRequest) *DeleteManagedOrganizationalUnitsInvoker {
	requestDef := GenReqDefForDeleteManagedOrganizationalUnits()
	return &DeleteManagedOrganizationalUnitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeregisterOrganizationalUnit 取消注册OU
//
// 将组织里的某个OU从RGC服务里取消注册。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DeregisterOrganizationalUnit(request *model.DeregisterOrganizationalUnitRequest) (*model.DeregisterOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForDeregisterOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeregisterOrganizationalUnitResponse), nil
	}
}

// DeregisterOrganizationalUnitInvoker 取消注册OU
func (c *RgcClient) DeregisterOrganizationalUnitInvoker(request *model.DeregisterOrganizationalUnitRequest) *DeregisterOrganizationalUnitInvoker {
	requestDef := GenReqDefForDeregisterOrganizationalUnit()
	return &DeregisterOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnrollAccount 纳管账号
//
// 将组织里的某个账号纳管到RGC服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) EnrollAccount(request *model.EnrollAccountRequest) (*model.EnrollAccountResponse, error) {
	requestDef := GenReqDefForEnrollAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnrollAccountResponse), nil
	}
}

// EnrollAccountInvoker 纳管账号
func (c *RgcClient) EnrollAccountInvoker(request *model.EnrollAccountRequest) *EnrollAccountInvoker {
	requestDef := GenReqDefForEnrollAccount()
	return &EnrollAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedAccounts 列举控制策略生效的纳管账号信息
//
// 列举控制策略生效的纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListManagedAccounts(request *model.ListManagedAccountsRequest) (*model.ListManagedAccountsResponse, error) {
	requestDef := GenReqDefForListManagedAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedAccountsResponse), nil
	}
}

// ListManagedAccountsInvoker 列举控制策略生效的纳管账号信息
func (c *RgcClient) ListManagedAccountsInvoker(request *model.ListManagedAccountsRequest) *ListManagedAccountsInvoker {
	requestDef := GenReqDefForListManagedAccounts()
	return &ListManagedAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedAccountsForParent 列出注册OU下的纳管账号信息
//
// 列出组织里某个注册OU下的所有纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListManagedAccountsForParent(request *model.ListManagedAccountsForParentRequest) (*model.ListManagedAccountsForParentResponse, error) {
	requestDef := GenReqDefForListManagedAccountsForParent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedAccountsForParentResponse), nil
	}
}

// ListManagedAccountsForParentInvoker 列出注册OU下的纳管账号信息
func (c *RgcClient) ListManagedAccountsForParentInvoker(request *model.ListManagedAccountsForParentRequest) *ListManagedAccountsForParentInvoker {
	requestDef := GenReqDefForListManagedAccountsForParent()
	return &ListManagedAccountsForParentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedOrganizationalUnits 列举控制策略生效的注册OU信息
//
// 列举控制策略生效的注册OU信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListManagedOrganizationalUnits(request *model.ListManagedOrganizationalUnitsRequest) (*model.ListManagedOrganizationalUnitsResponse, error) {
	requestDef := GenReqDefForListManagedOrganizationalUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedOrganizationalUnitsResponse), nil
	}
}

// ListManagedOrganizationalUnitsInvoker 列举控制策略生效的注册OU信息
func (c *RgcClient) ListManagedOrganizationalUnitsInvoker(request *model.ListManagedOrganizationalUnitsRequest) *ListManagedOrganizationalUnitsInvoker {
	requestDef := GenReqDefForListManagedOrganizationalUnits()
	return &ListManagedOrganizationalUnitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOperation 查询已注册OU和纳管账号操作过程信息列表
//
// 查询在RGC服务里已注册OU和纳管账号操作的过程信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListOperation(request *model.ListOperationRequest) (*model.ListOperationResponse, error) {
	requestDef := GenReqDefForListOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOperationResponse), nil
	}
}

// ListOperationInvoker 查询已注册OU和纳管账号操作过程信息列表
func (c *RgcClient) ListOperationInvoker(request *model.ListOperationRequest) *ListOperationInvoker {
	requestDef := GenReqDefForListOperation()
	return &ListOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReRegisterOrganizationalUnit 重新注册OU
//
// 重新注册组织里的某个OU到RGC服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ReRegisterOrganizationalUnit(request *model.ReRegisterOrganizationalUnitRequest) (*model.ReRegisterOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForReRegisterOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReRegisterOrganizationalUnitResponse), nil
	}
}

// ReRegisterOrganizationalUnitInvoker 重新注册OU
func (c *RgcClient) ReRegisterOrganizationalUnitInvoker(request *model.ReRegisterOrganizationalUnitRequest) *ReRegisterOrganizationalUnitInvoker {
	requestDef := GenReqDefForReRegisterOrganizationalUnit()
	return &ReRegisterOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterOrganizationalUnit 注册OU
//
// 将组织里的某个OU注册到RGC服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) RegisterOrganizationalUnit(request *model.RegisterOrganizationalUnitRequest) (*model.RegisterOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForRegisterOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterOrganizationalUnitResponse), nil
	}
}

// RegisterOrganizationalUnitInvoker 注册OU
func (c *RgcClient) RegisterOrganizationalUnitInvoker(request *model.RegisterOrganizationalUnitRequest) *RegisterOrganizationalUnitInvoker {
	requestDef := GenReqDefForRegisterOrganizationalUnit()
	return &RegisterOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedAccount 查询纳管账号信息
//
// 查询组织里某个纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedAccount(request *model.ShowManagedAccountRequest) (*model.ShowManagedAccountResponse, error) {
	requestDef := GenReqDefForShowManagedAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedAccountResponse), nil
	}
}

// ShowManagedAccountInvoker 查询纳管账号信息
func (c *RgcClient) ShowManagedAccountInvoker(request *model.ShowManagedAccountRequest) *ShowManagedAccountInvoker {
	requestDef := GenReqDefForShowManagedAccount()
	return &ShowManagedAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedAccountTemplate 查询纳管账号的模板信息
//
// 查询组织里某个纳管账号的模板信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedAccountTemplate(request *model.ShowManagedAccountTemplateRequest) (*model.ShowManagedAccountTemplateResponse, error) {
	requestDef := GenReqDefForShowManagedAccountTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedAccountTemplateResponse), nil
	}
}

// ShowManagedAccountTemplateInvoker 查询纳管账号的模板信息
func (c *RgcClient) ShowManagedAccountTemplateInvoker(request *model.ShowManagedAccountTemplateRequest) *ShowManagedAccountTemplateInvoker {
	requestDef := GenReqDefForShowManagedAccountTemplate()
	return &ShowManagedAccountTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedCoreAccount 列出核心纳管账号
//
// 列出组织里的所有核心纳管账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedCoreAccount(request *model.ShowManagedCoreAccountRequest) (*model.ShowManagedCoreAccountResponse, error) {
	requestDef := GenReqDefForShowManagedCoreAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedCoreAccountResponse), nil
	}
}

// ShowManagedCoreAccountInvoker 列出核心纳管账号
func (c *RgcClient) ShowManagedCoreAccountInvoker(request *model.ShowManagedCoreAccountRequest) *ShowManagedCoreAccountInvoker {
	requestDef := GenReqDefForShowManagedCoreAccount()
	return &ShowManagedCoreAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowManagedOrganizationalUnit 查询已注册OU信息
//
// 查询在RGC服务里的注册OU信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowManagedOrganizationalUnit(request *model.ShowManagedOrganizationalUnitRequest) (*model.ShowManagedOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForShowManagedOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowManagedOrganizationalUnitResponse), nil
	}
}

// ShowManagedOrganizationalUnitInvoker 查询已注册OU信息
func (c *RgcClient) ShowManagedOrganizationalUnitInvoker(request *model.ShowManagedOrganizationalUnitRequest) *ShowManagedOrganizationalUnitInvoker {
	requestDef := GenReqDefForShowManagedOrganizationalUnit()
	return &ShowManagedOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOperation 查询注册过程信息
//
// 查询在RGC服务里注册/取消注册的过程信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowOperation(request *model.ShowOperationRequest) (*model.ShowOperationResponse, error) {
	requestDef := GenReqDefForShowOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOperationResponse), nil
	}
}

// ShowOperationInvoker 查询注册过程信息
func (c *RgcClient) ShowOperationInvoker(request *model.ShowOperationRequest) *ShowOperationInvoker {
	requestDef := GenReqDefForShowOperation()
	return &ShowOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnEnrollAccount 取消纳管账号
//
// 将组织里的某个账号从RGC服务里取消纳管。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) UnEnrollAccount(request *model.UnEnrollAccountRequest) (*model.UnEnrollAccountResponse, error) {
	requestDef := GenReqDefForUnEnrollAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnEnrollAccountResponse), nil
	}
}

// UnEnrollAccountInvoker 取消纳管账号
func (c *RgcClient) UnEnrollAccountInvoker(request *model.UnEnrollAccountRequest) *UnEnrollAccountInvoker {
	requestDef := GenReqDefForUnEnrollAccount()
	return &UnEnrollAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateManagedAccount 更新纳管账号
//
// 更新组织里某个已在RGC服务的纳管账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) UpdateManagedAccount(request *model.UpdateManagedAccountRequest) (*model.UpdateManagedAccountResponse, error) {
	requestDef := GenReqDefForUpdateManagedAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateManagedAccountResponse), nil
	}
}

// UpdateManagedAccountInvoker 更新纳管账号
func (c *RgcClient) UpdateManagedAccountInvoker(request *model.UpdateManagedAccountRequest) *UpdateManagedAccountInvoker {
	requestDef := GenReqDefForUpdateManagedAccount()
	return &UpdateManagedAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplate 创建模板
//
// 创建RFS模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// CreateTemplateInvoker 创建模板
func (c *RgcClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplate 删除模板
//
// 删除RFS模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// DeleteTemplateInvoker 删除模板
func (c *RgcClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPredefinedTemplates 查询预置模板列表
//
// 查询预置模板列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ListPredefinedTemplates(request *model.ListPredefinedTemplatesRequest) (*model.ListPredefinedTemplatesResponse, error) {
	requestDef := GenReqDefForListPredefinedTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPredefinedTemplatesResponse), nil
	}
}

// ListPredefinedTemplatesInvoker 查询预置模板列表
func (c *RgcClient) ListPredefinedTemplatesInvoker(request *model.ListPredefinedTemplatesRequest) *ListPredefinedTemplatesInvoker {
	requestDef := GenReqDefForListPredefinedTemplates()
	return &ListPredefinedTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplateDeployParams 查询模板的部署参数
//
// 查询模板的部署参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RgcClient) ShowTemplateDeployParams(request *model.ShowTemplateDeployParamsRequest) (*model.ShowTemplateDeployParamsResponse, error) {
	requestDef := GenReqDefForShowTemplateDeployParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDeployParamsResponse), nil
	}
}

// ShowTemplateDeployParamsInvoker 查询模板的部署参数
func (c *RgcClient) ShowTemplateDeployParamsInvoker(request *model.ShowTemplateDeployParamsRequest) *ShowTemplateDeployParamsInvoker {
	requestDef := GenReqDefForShowTemplateDeployParams()
	return &ShowTemplateDeployParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
