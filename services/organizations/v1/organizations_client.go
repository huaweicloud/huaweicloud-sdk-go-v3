package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/organizations/v1/model"
)

type OrganizationsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewOrganizationsClient(hcClient *httpclient.HcHttpClient) *OrganizationsClient {
	return &OrganizationsClient{HcClient: hcClient}
}

func OrganizationsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CloseAccount 关闭账号
//
// 关闭账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) CloseAccount(request *model.CloseAccountRequest) (*model.CloseAccountResponse, error) {
	requestDef := GenReqDefForCloseAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CloseAccountResponse), nil
	}
}

// CloseAccountInvoker 关闭账号
func (c *OrganizationsClient) CloseAccountInvoker(request *model.CloseAccountRequest) *CloseAccountInvoker {
	requestDef := GenReqDefForCloseAccount()
	return &CloseAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccount 创建账号
//
// 创建一个账号，生成的账号将自动成为调用此接口的账号所属组织的成员。此操作只能由组织的管理账号调用。组织云服务将在新账号中创建所需的服务关联委托和账号访问委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) CreateAccount(request *model.CreateAccountRequest) (*model.CreateAccountResponse, error) {
	requestDef := GenReqDefForCreateAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccountResponse), nil
	}
}

// CreateAccountInvoker 创建账号
func (c *OrganizationsClient) CreateAccountInvoker(request *model.CreateAccountRequest) *CreateAccountInvoker {
	requestDef := GenReqDefForCreateAccount()
	return &CreateAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteAccount 邀请账号加入组织
//
// 向另一个账号发送邀请，受邀账号将以成员账号加入您的组织。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) InviteAccount(request *model.InviteAccountRequest) (*model.InviteAccountResponse, error) {
	requestDef := GenReqDefForInviteAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteAccountResponse), nil
	}
}

// InviteAccountInvoker 邀请账号加入组织
func (c *OrganizationsClient) InviteAccountInvoker(request *model.InviteAccountRequest) *InviteAccountInvoker {
	requestDef := GenReqDefForInviteAccount()
	return &InviteAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccounts 列出组织中的账号
//
// 列出组织中的账号。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。如果指定父级组织单元，则将获得作为父级直系子级的所有账号的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListAccounts(request *model.ListAccountsRequest) (*model.ListAccountsResponse, error) {
	requestDef := GenReqDefForListAccounts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountsResponse), nil
	}
}

// ListAccountsInvoker 列出组织中的账号
func (c *OrganizationsClient) ListAccountsInvoker(request *model.ListAccountsRequest) *ListAccountsInvoker {
	requestDef := GenReqDefForListAccounts()
	return &ListAccountsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloseAccountStatuses 列出关闭账号的状态
//
// 列出组织中指定状态的账号关闭请求。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListCloseAccountStatuses(request *model.ListCloseAccountStatusesRequest) (*model.ListCloseAccountStatusesResponse, error) {
	requestDef := GenReqDefForListCloseAccountStatuses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloseAccountStatusesResponse), nil
	}
}

// ListCloseAccountStatusesInvoker 列出关闭账号的状态
func (c *OrganizationsClient) ListCloseAccountStatusesInvoker(request *model.ListCloseAccountStatusesRequest) *ListCloseAccountStatusesInvoker {
	requestDef := GenReqDefForListCloseAccountStatuses()
	return &ListCloseAccountStatusesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCreateAccountStatuses 列出创建账号的状态
//
// 列出组织中指定状态的账号创建请求。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListCreateAccountStatuses(request *model.ListCreateAccountStatusesRequest) (*model.ListCreateAccountStatusesResponse, error) {
	requestDef := GenReqDefForListCreateAccountStatuses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCreateAccountStatusesResponse), nil
	}
}

// ListCreateAccountStatusesInvoker 列出创建账号的状态
func (c *OrganizationsClient) ListCreateAccountStatusesInvoker(request *model.ListCreateAccountStatusesRequest) *ListCreateAccountStatusesInvoker {
	requestDef := GenReqDefForListCreateAccountStatuses()
	return &ListCreateAccountStatusesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MoveAccount 移动账号
//
// 将账号从其当前源位置（根或组织单元）移动到指定的目标位置（根或组织单元）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) MoveAccount(request *model.MoveAccountRequest) (*model.MoveAccountResponse, error) {
	requestDef := GenReqDefForMoveAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MoveAccountResponse), nil
	}
}

// MoveAccountInvoker 移动账号
func (c *OrganizationsClient) MoveAccountInvoker(request *model.MoveAccountRequest) *MoveAccountInvoker {
	requestDef := GenReqDefForMoveAccount()
	return &MoveAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveAccount 移除指定的账号
//
// 从组织中移除指定的账号。移除的账号将成为一个独立账号，该账号不是任何组织的成员。此操作只能由组织的管理账号调用。只有当账号配置了作为独立账号运行所需的信息时，您才能从组织中移除账号。注意，要移除的账号不能是组织启用的任何服务的委托管理员账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) RemoveAccount(request *model.RemoveAccountRequest) (*model.RemoveAccountResponse, error) {
	requestDef := GenReqDefForRemoveAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveAccountResponse), nil
	}
}

// RemoveAccountInvoker 移除指定的账号
func (c *OrganizationsClient) RemoveAccountInvoker(request *model.RemoveAccountRequest) *RemoveAccountInvoker {
	requestDef := GenReqDefForRemoveAccount()
	return &RemoveAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccount 查询账号信息
//
// 查询有关指定账号的信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowAccount(request *model.ShowAccountRequest) (*model.ShowAccountResponse, error) {
	requestDef := GenReqDefForShowAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccountResponse), nil
	}
}

// ShowAccountInvoker 查询账号信息
func (c *OrganizationsClient) ShowAccountInvoker(request *model.ShowAccountRequest) *ShowAccountInvoker {
	requestDef := GenReqDefForShowAccount()
	return &ShowAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCreateAccountStatus 查询有关创建账号状态的信息
//
// 检索创建账号的异步请求的当前状态。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowCreateAccountStatus(request *model.ShowCreateAccountStatusRequest) (*model.ShowCreateAccountStatusResponse, error) {
	requestDef := GenReqDefForShowCreateAccountStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCreateAccountStatusResponse), nil
	}
}

// ShowCreateAccountStatusInvoker 查询有关创建账号状态的信息
func (c *OrganizationsClient) ShowCreateAccountStatusInvoker(request *model.ShowCreateAccountStatusRequest) *ShowCreateAccountStatusInvoker {
	requestDef := GenReqDefForShowCreateAccountStatus()
	return &ShowCreateAccountStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccount 更新账号信息
//
// 更新指定的账号信息。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) UpdateAccount(request *model.UpdateAccountRequest) (*model.UpdateAccountResponse, error) {
	requestDef := GenReqDefForUpdateAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccountResponse), nil
	}
}

// UpdateAccountInvoker 更新账号信息
func (c *OrganizationsClient) UpdateAccountInvoker(request *model.UpdateAccountRequest) *UpdateAccountInvoker {
	requestDef := GenReqDefForUpdateAccount()
	return &UpdateAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeregisterDelegatedAdministrator 注销服务的委托管理员
//
// 删除指定成员账号作为指定服务的委托管理员。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DeregisterDelegatedAdministrator(request *model.DeregisterDelegatedAdministratorRequest) (*model.DeregisterDelegatedAdministratorResponse, error) {
	requestDef := GenReqDefForDeregisterDelegatedAdministrator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeregisterDelegatedAdministratorResponse), nil
	}
}

// DeregisterDelegatedAdministratorInvoker 注销服务的委托管理员
func (c *OrganizationsClient) DeregisterDelegatedAdministratorInvoker(request *model.DeregisterDelegatedAdministratorRequest) *DeregisterDelegatedAdministratorInvoker {
	requestDef := GenReqDefForDeregisterDelegatedAdministrator()
	return &DeregisterDelegatedAdministratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDelegatedAdministrators 列出此组织中指定为委托管理员的账号
//
// 列出在此组织中指定为委派管理员的账号。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListDelegatedAdministrators(request *model.ListDelegatedAdministratorsRequest) (*model.ListDelegatedAdministratorsResponse, error) {
	requestDef := GenReqDefForListDelegatedAdministrators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDelegatedAdministratorsResponse), nil
	}
}

// ListDelegatedAdministratorsInvoker 列出此组织中指定为委托管理员的账号
func (c *OrganizationsClient) ListDelegatedAdministratorsInvoker(request *model.ListDelegatedAdministratorsRequest) *ListDelegatedAdministratorsInvoker {
	requestDef := GenReqDefForListDelegatedAdministrators()
	return &ListDelegatedAdministratorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDelegatedServices 列出指定账号是其委托管理员的服务
//
// 列出指定账号是其委托管理员的服务。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListDelegatedServices(request *model.ListDelegatedServicesRequest) (*model.ListDelegatedServicesResponse, error) {
	requestDef := GenReqDefForListDelegatedServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDelegatedServicesResponse), nil
	}
}

// ListDelegatedServicesInvoker 列出指定账号是其委托管理员的服务
func (c *OrganizationsClient) ListDelegatedServicesInvoker(request *model.ListDelegatedServicesRequest) *ListDelegatedServicesInvoker {
	requestDef := GenReqDefForListDelegatedServices()
	return &ListDelegatedServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterDelegatedAdministrator 注册作为服务委托管理员
//
// 指定成员账号能够管理指定服务的组织功能。此接口授予委托管理员对组织服务数据的只读访问权限。委托管理员账号中的IAM用户仍需要IAM权限才能访问和管理服务。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) RegisterDelegatedAdministrator(request *model.RegisterDelegatedAdministratorRequest) (*model.RegisterDelegatedAdministratorResponse, error) {
	requestDef := GenReqDefForRegisterDelegatedAdministrator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterDelegatedAdministratorResponse), nil
	}
}

// RegisterDelegatedAdministratorInvoker 注册作为服务委托管理员
func (c *OrganizationsClient) RegisterDelegatedAdministratorInvoker(request *model.RegisterDelegatedAdministratorRequest) *RegisterDelegatedAdministratorInvoker {
	requestDef := GenReqDefForRegisterDelegatedAdministrator()
	return &RegisterDelegatedAdministratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AcceptHandshake 接受邀请
//
// 向邀请的发起方发送应答，接受加入组织邀请。在您接受邀请后，此邀请信息将继续保留并出现在相关API的返回结果中，保留期限为30天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) AcceptHandshake(request *model.AcceptHandshakeRequest) (*model.AcceptHandshakeResponse, error) {
	requestDef := GenReqDefForAcceptHandshake()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptHandshakeResponse), nil
	}
}

// AcceptHandshakeInvoker 接受邀请
func (c *OrganizationsClient) AcceptHandshakeInvoker(request *model.AcceptHandshakeRequest) *AcceptHandshakeInvoker {
	requestDef := GenReqDefForAcceptHandshake()
	return &AcceptHandshakeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelHandshake 取消邀请
//
// 取消邀请，此时邀请状态将设置为已取消。此接口只能由发起邀请的账号调用。取消邀请后，此邀请信息将继续保留并出现在相关API的返回结果中，保留期限为30天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) CancelHandshake(request *model.CancelHandshakeRequest) (*model.CancelHandshakeResponse, error) {
	requestDef := GenReqDefForCancelHandshake()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelHandshakeResponse), nil
	}
}

// CancelHandshakeInvoker 取消邀请
func (c *OrganizationsClient) CancelHandshakeInvoker(request *model.CancelHandshakeRequest) *CancelHandshakeInvoker {
	requestDef := GenReqDefForCancelHandshake()
	return &CancelHandshakeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeclineHandshake 拒绝邀请
//
// 拒绝邀请请求。受邀账号拒绝邀请，此时当前邀请状态将设置为拒绝，邀请停止。此接口只能由受邀账号调用。邀请发起者无法再次激活被拒绝的邀请，但可以重新发送新的邀请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DeclineHandshake(request *model.DeclineHandshakeRequest) (*model.DeclineHandshakeResponse, error) {
	requestDef := GenReqDefForDeclineHandshake()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeclineHandshakeResponse), nil
	}
}

// DeclineHandshakeInvoker 拒绝邀请
func (c *OrganizationsClient) DeclineHandshakeInvoker(request *model.DeclineHandshakeRequest) *DeclineHandshakeInvoker {
	requestDef := GenReqDefForDeclineHandshake()
	return &DeclineHandshakeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHandshakes 列出发送的邀请
//
// 列出所属组织发送的邀请。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListHandshakes(request *model.ListHandshakesRequest) (*model.ListHandshakesResponse, error) {
	requestDef := GenReqDefForListHandshakes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHandshakesResponse), nil
	}
}

// ListHandshakesInvoker 列出发送的邀请
func (c *OrganizationsClient) ListHandshakesInvoker(request *model.ListHandshakesRequest) *ListHandshakesInvoker {
	requestDef := GenReqDefForListHandshakes()
	return &ListHandshakesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReceivedHandshakes 列出收到的邀请
//
// 列出账号收到的所有邀请。此操作可以由任何账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListReceivedHandshakes(request *model.ListReceivedHandshakesRequest) (*model.ListReceivedHandshakesResponse, error) {
	requestDef := GenReqDefForListReceivedHandshakes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReceivedHandshakesResponse), nil
	}
}

// ListReceivedHandshakesInvoker 列出收到的邀请
func (c *OrganizationsClient) ListReceivedHandshakesInvoker(request *model.ListReceivedHandshakesRequest) *ListReceivedHandshakesInvoker {
	requestDef := GenReqDefForListReceivedHandshakes()
	return &ListReceivedHandshakesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHandshake 查询邀请相关信息
//
// 查询组织中已有账号邀请的相关信息。此接口可以由组织中的任何账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowHandshake(request *model.ShowHandshakeRequest) (*model.ShowHandshakeResponse, error) {
	requestDef := GenReqDefForShowHandshake()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHandshakeResponse), nil
	}
}

// ShowHandshakeInvoker 查询邀请相关信息
func (c *OrganizationsClient) ShowHandshakeInvoker(request *model.ShowHandshakeRequest) *ShowHandshakeInvoker {
	requestDef := GenReqDefForShowHandshake()
	return &ShowHandshakeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEntities 列出组织中的根、组织单元和账号
//
// 列出组织中包含的所有根、组织单元和账号。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。您可以通过指定父ID和子ID参数来过滤实体。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListEntities(request *model.ListEntitiesRequest) (*model.ListEntitiesResponse, error) {
	requestDef := GenReqDefForListEntities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEntitiesResponse), nil
	}
}

// ListEntitiesInvoker 列出组织中的根、组织单元和账号
func (c *OrganizationsClient) ListEntitiesInvoker(request *model.ListEntitiesRequest) *ListEntitiesInvoker {
	requestDef := GenReqDefForListEntities()
	return &ListEntitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 列出租户的组织配额
//
// 列出租户的组织配额。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 列出租户的组织配额
func (c *OrganizationsClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServices 列出所有可以与组织服务集成的云服务
//
// 列出所有可以与组织服务集成的云服务。集成后云服务将成为组织的可信服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListServices(request *model.ListServicesRequest) (*model.ListServicesResponse, error) {
	requestDef := GenReqDefForListServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicesResponse), nil
	}
}

// ListServicesInvoker 列出所有可以与组织服务集成的云服务
func (c *OrganizationsClient) ListServicesInvoker(request *model.ListServicesRequest) *ListServicesInvoker {
	requestDef := GenReqDefForListServices()
	return &ListServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagPolicyServices 列出被添加到标签策略强制执行的资源类型
//
// 列出被添加到标签策略强制执行的资源类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListTagPolicyServices(request *model.ListTagPolicyServicesRequest) (*model.ListTagPolicyServicesResponse, error) {
	requestDef := GenReqDefForListTagPolicyServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagPolicyServicesResponse), nil
	}
}

// ListTagPolicyServicesInvoker 列出被添加到标签策略强制执行的资源类型
func (c *OrganizationsClient) ListTagPolicyServicesInvoker(request *model.ListTagPolicyServicesRequest) *ListTagPolicyServicesInvoker {
	requestDef := GenReqDefForListTagPolicyServices()
	return &ListTagPolicyServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEffectivePolicies 查询有效的策略
//
// 查询指定策略类型和账号的有效策略信息。当前此接口不支持查询服务控制策略（service_control_policy）。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowEffectivePolicies(request *model.ShowEffectivePoliciesRequest) (*model.ShowEffectivePoliciesResponse, error) {
	requestDef := GenReqDefForShowEffectivePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEffectivePoliciesResponse), nil
	}
}

// ShowEffectivePoliciesInvoker 查询有效的策略
func (c *OrganizationsClient) ShowEffectivePoliciesInvoker(request *model.ShowEffectivePoliciesRequest) *ShowEffectivePoliciesInvoker {
	requestDef := GenReqDefForShowEffectivePolicies()
	return &ShowEffectivePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrganization 创建组织
//
// 创建组织。调用此接口的账号将自动成为新组织的管理账号，调用此接口的账号凭证必须是新组织管理账号的账号凭证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) CreateOrganization(request *model.CreateOrganizationRequest) (*model.CreateOrganizationResponse, error) {
	requestDef := GenReqDefForCreateOrganization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrganizationResponse), nil
	}
}

// CreateOrganizationInvoker 创建组织
func (c *OrganizationsClient) CreateOrganizationInvoker(request *model.CreateOrganizationRequest) *CreateOrganizationInvoker {
	requestDef := GenReqDefForCreateOrganization()
	return &CreateOrganizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOrganization 删除组织
//
// 删除组织。您必须使用管理账号才能删除组织，并且先移除组织中的所有账号、组织单元和策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DeleteOrganization(request *model.DeleteOrganizationRequest) (*model.DeleteOrganizationResponse, error) {
	requestDef := GenReqDefForDeleteOrganization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOrganizationResponse), nil
	}
}

// DeleteOrganizationInvoker 删除组织
func (c *OrganizationsClient) DeleteOrganizationInvoker(request *model.DeleteOrganizationRequest) *DeleteOrganizationInvoker {
	requestDef := GenReqDefForDeleteOrganization()
	return &DeleteOrganizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LeaveOrganization 离开当前组织
//
// 此操作只能由组织的成员账号调用。只有当组织账号配置了作为独立账号运行所需的信息时，您才能作为成员账号离开组织。要离开的账号不能是组织启用的任何服务的委托管理员账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) LeaveOrganization(request *model.LeaveOrganizationRequest) (*model.LeaveOrganizationResponse, error) {
	requestDef := GenReqDefForLeaveOrganization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LeaveOrganizationResponse), nil
	}
}

// LeaveOrganizationInvoker 离开当前组织
func (c *OrganizationsClient) LeaveOrganizationInvoker(request *model.LeaveOrganizationRequest) *LeaveOrganizationInvoker {
	requestDef := GenReqDefForLeaveOrganization()
	return &LeaveOrganizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRoots 列出组织的根
//
// 列出当前组织的根。此接口只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListRoots(request *model.ListRootsRequest) (*model.ListRootsResponse, error) {
	requestDef := GenReqDefForListRoots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRootsResponse), nil
	}
}

// ListRootsInvoker 列出组织的根
func (c *OrganizationsClient) ListRootsInvoker(request *model.ListRootsRequest) *ListRootsInvoker {
	requestDef := GenReqDefForListRoots()
	return &ListRootsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganization 查询所属组织信息
//
// 查询账号所属组织的信息。此操作可以由组织中的任何账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowOrganization(request *model.ShowOrganizationRequest) (*model.ShowOrganizationResponse, error) {
	requestDef := GenReqDefForShowOrganization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationResponse), nil
	}
}

// ShowOrganizationInvoker 查询所属组织信息
func (c *OrganizationsClient) ShowOrganizationInvoker(request *model.ShowOrganizationRequest) *ShowOrganizationInvoker {
	requestDef := GenReqDefForShowOrganization()
	return &ShowOrganizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrganizationalUnit 创建组织单元
//
// 在根或父组织单元中创建组织单元。组织单元是账号的容器，使您能够对账号进行分组管理，并根据业务要求应用策略。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) CreateOrganizationalUnit(request *model.CreateOrganizationalUnitRequest) (*model.CreateOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForCreateOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrganizationalUnitResponse), nil
	}
}

// CreateOrganizationalUnitInvoker 创建组织单元
func (c *OrganizationsClient) CreateOrganizationalUnitInvoker(request *model.CreateOrganizationalUnitRequest) *CreateOrganizationalUnitInvoker {
	requestDef := GenReqDefForCreateOrganizationalUnit()
	return &CreateOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOrganizationalUnit 删除组织单元
//
// 从根或其他组织单元中删除组织单元。前提是您必须先移除该组织单元中的所有成员账号或将成员账号移动至其他组织单元，必须删除该组织单元中的所有子组织单元。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DeleteOrganizationalUnit(request *model.DeleteOrganizationalUnitRequest) (*model.DeleteOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForDeleteOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOrganizationalUnitResponse), nil
	}
}

// DeleteOrganizationalUnitInvoker 删除组织单元
func (c *OrganizationsClient) DeleteOrganizationalUnitInvoker(request *model.DeleteOrganizationalUnitRequest) *DeleteOrganizationalUnitInvoker {
	requestDef := GenReqDefForDeleteOrganizationalUnit()
	return &DeleteOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrganizationalUnits 列出组织单元
//
// 列出组织中的所有组织单元。如果指定父级组织单元，则将获得父级直系子级的组织单元列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListOrganizationalUnits(request *model.ListOrganizationalUnitsRequest) (*model.ListOrganizationalUnitsResponse, error) {
	requestDef := GenReqDefForListOrganizationalUnits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrganizationalUnitsResponse), nil
	}
}

// ListOrganizationalUnitsInvoker 列出组织单元
func (c *OrganizationsClient) ListOrganizationalUnitsInvoker(request *model.ListOrganizationalUnitsRequest) *ListOrganizationalUnitsInvoker {
	requestDef := GenReqDefForListOrganizationalUnits()
	return &ListOrganizationalUnitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationalUnit 查询有关组织单元的信息
//
// 查询有关组织单元的信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowOrganizationalUnit(request *model.ShowOrganizationalUnitRequest) (*model.ShowOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForShowOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationalUnitResponse), nil
	}
}

// ShowOrganizationalUnitInvoker 查询有关组织单元的信息
func (c *OrganizationsClient) ShowOrganizationalUnitInvoker(request *model.ShowOrganizationalUnitRequest) *ShowOrganizationalUnitInvoker {
	requestDef := GenReqDefForShowOrganizationalUnit()
	return &ShowOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateOrganizationalUnit 更改组织单元名称
//
// 重命名指定的组织单元。重命名后组织单元ID不变，下属子组织单元和下属账号不变，组织单元绑定的策略不变。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) UpdateOrganizationalUnit(request *model.UpdateOrganizationalUnitRequest) (*model.UpdateOrganizationalUnitResponse, error) {
	requestDef := GenReqDefForUpdateOrganizationalUnit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOrganizationalUnitResponse), nil
	}
}

// UpdateOrganizationalUnitInvoker 更改组织单元名称
func (c *OrganizationsClient) UpdateOrganizationalUnitInvoker(request *model.UpdateOrganizationalUnitRequest) *UpdateOrganizationalUnitInvoker {
	requestDef := GenReqDefForUpdateOrganizationalUnit()
	return &UpdateOrganizationalUnitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachPolicy 将策略跟实体绑定
//
// 绑定策略到根、组织单元或个人账号。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) AttachPolicy(request *model.AttachPolicyRequest) (*model.AttachPolicyResponse, error) {
	requestDef := GenReqDefForAttachPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachPolicyResponse), nil
	}
}

// AttachPolicyInvoker 将策略跟实体绑定
func (c *OrganizationsClient) AttachPolicyInvoker(request *model.AttachPolicyRequest) *AttachPolicyInvoker {
	requestDef := GenReqDefForAttachPolicy()
	return &AttachPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicy 创建策略
//
// 创建指定类型的策略。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyResponse), nil
	}
}

// CreatePolicyInvoker 创建策略
func (c *OrganizationsClient) CreatePolicyInvoker(request *model.CreatePolicyRequest) *CreatePolicyInvoker {
	requestDef := GenReqDefForCreatePolicy()
	return &CreatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除策略
//
// 从组织中删除指定的策略。在执行此操作之前，必须首先将策略跟所有组织单元、根和账号解绑。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

// DeletePolicyInvoker 删除策略
func (c *OrganizationsClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachPolicy 将策略跟实体解绑
//
// 从根、组织单元或账号解绑策略。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DetachPolicy(request *model.DetachPolicyRequest) (*model.DetachPolicyResponse, error) {
	requestDef := GenReqDefForDetachPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachPolicyResponse), nil
	}
}

// DetachPolicyInvoker 将策略跟实体解绑
func (c *OrganizationsClient) DetachPolicyInvoker(request *model.DetachPolicyRequest) *DetachPolicyInvoker {
	requestDef := GenReqDefForDetachPolicy()
	return &DetachPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisablePolicyType 禁用根中的策略类型
//
// 禁用根中的策略类型。只有在根中启用了特定类型的策略，才能将该类型的策略绑定到根中的实体。执行此操作后，您不能再将指定类型的策略绑定到该根或该根中的任何组织单元或账号。这是在后台执行的异步请求。您可以使用ListRoots查看指定根的策略类型的状态。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DisablePolicyType(request *model.DisablePolicyTypeRequest) (*model.DisablePolicyTypeResponse, error) {
	requestDef := GenReqDefForDisablePolicyType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisablePolicyTypeResponse), nil
	}
}

// DisablePolicyTypeInvoker 禁用根中的策略类型
func (c *OrganizationsClient) DisablePolicyTypeInvoker(request *model.DisablePolicyTypeRequest) *DisablePolicyTypeInvoker {
	requestDef := GenReqDefForDisablePolicyType()
	return &DisablePolicyTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnablePolicyType 在根中启用策略类型
//
// 在根中启用策略类型。在根中启用策略类型后，您可以将该类型的策略绑定到根、或该根中的任何组织单元和账号。这是在后台执行的异步请求。您可以使用ListRoots查看指定根的策略类型的状态。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) EnablePolicyType(request *model.EnablePolicyTypeRequest) (*model.EnablePolicyTypeResponse, error) {
	requestDef := GenReqDefForEnablePolicyType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnablePolicyTypeResponse), nil
	}
}

// EnablePolicyTypeInvoker 在根中启用策略类型
func (c *OrganizationsClient) EnablePolicyTypeInvoker(request *model.EnablePolicyTypeRequest) *EnablePolicyTypeInvoker {
	requestDef := GenReqDefForEnablePolicyType()
	return &EnablePolicyTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEntitiesForPolicy 列出跟指定策略绑定的所有实体
//
// 列出跟指定策略绑定的所有根、组织单元和账号。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListEntitiesForPolicy(request *model.ListEntitiesForPolicyRequest) (*model.ListEntitiesForPolicyResponse, error) {
	requestDef := GenReqDefForListEntitiesForPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEntitiesForPolicyResponse), nil
	}
}

// ListEntitiesForPolicyInvoker 列出跟指定策略绑定的所有实体
func (c *OrganizationsClient) ListEntitiesForPolicyInvoker(request *model.ListEntitiesForPolicyRequest) *ListEntitiesForPolicyInvoker {
	requestDef := GenReqDefForListEntitiesForPolicy()
	return &ListEntitiesForPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicies 列出策略
//
// 列出组织中的所有策略。如果指定了资源ID，例如组织单元ID或账号ID，则将获得该资源已绑定的策略列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListPolicies(request *model.ListPoliciesRequest) (*model.ListPoliciesResponse, error) {
	requestDef := GenReqDefForListPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoliciesResponse), nil
	}
}

// ListPoliciesInvoker 列出策略
func (c *OrganizationsClient) ListPoliciesInvoker(request *model.ListPoliciesRequest) *ListPoliciesInvoker {
	requestDef := GenReqDefForListPolicies()
	return &ListPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicy 查询策略相关信息
//
// 检索策略的相关信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyResponse), nil
	}
}

// ShowPolicyInvoker 查询策略相关信息
func (c *OrganizationsClient) ShowPolicyInvoker(request *model.ShowPolicyRequest) *ShowPolicyInvoker {
	requestDef := GenReqDefForShowPolicy()
	return &ShowPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicy 更新策略
//
// 更新策略，可以更新策略的名称、描述或内容。如果不提供任何参数，则策略将保持不变。您不能更改策略的类型。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyResponse), nil
	}
}

// UpdatePolicyInvoker 更新策略
func (c *OrganizationsClient) UpdatePolicyInvoker(request *model.UpdatePolicyRequest) *UpdatePolicyInvoker {
	requestDef := GenReqDefForUpdatePolicy()
	return &UpdatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTagResource 为指定资源类型添加标签
//
// 向指定的资源类型添加一个或多个标签。目前，您可以将标签附加到组织中的账号、组织单元、根和策略。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) CreateTagResource(request *model.CreateTagResourceRequest) (*model.CreateTagResourceResponse, error) {
	requestDef := GenReqDefForCreateTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResourceResponse), nil
	}
}

// CreateTagResourceInvoker 为指定资源类型添加标签
func (c *OrganizationsClient) CreateTagResourceInvoker(request *model.CreateTagResourceRequest) *CreateTagResourceInvoker {
	requestDef := GenReqDefForCreateTagResource()
	return &CreateTagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTagResource 从指定资源类型中删除指定主键标签
//
// 从指定资源类型中删除具有指定主键的任何标签。您可以将标签绑定到组织中的账号、组织单元、根和策略。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DeleteTagResource(request *model.DeleteTagResourceRequest) (*model.DeleteTagResourceResponse, error) {
	requestDef := GenReqDefForDeleteTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResourceResponse), nil
	}
}

// DeleteTagResourceInvoker 从指定资源类型中删除指定主键标签
func (c *OrganizationsClient) DeleteTagResourceInvoker(request *model.DeleteTagResourceRequest) *DeleteTagResourceInvoker {
	requestDef := GenReqDefForDeleteTagResource()
	return &DeleteTagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstances 根据资源类型及标签信息查询实例列表
//
// 根据资源类型及标签信息查询实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

// ListResourceInstancesInvoker 根据资源类型及标签信息查询实例列表
func (c *OrganizationsClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTags 查询资源标签
//
// 查询资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListResourceTags(request *model.ListResourceTagsRequest) (*model.ListResourceTagsResponse, error) {
	requestDef := GenReqDefForListResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsResponse), nil
	}
}

// ListResourceTagsInvoker 查询资源标签
func (c *OrganizationsClient) ListResourceTagsInvoker(request *model.ListResourceTagsRequest) *ListResourceTagsInvoker {
	requestDef := GenReqDefForListResourceTags()
	return &ListResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagResources 列出绑定到指定资源类型的标签
//
// 列出绑定到指定资源类型的标签。您可以将标签附加到组织中的账号、组织单元、根和策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListTagResources(request *model.ListTagResourcesRequest) (*model.ListTagResourcesResponse, error) {
	requestDef := GenReqDefForListTagResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagResourcesResponse), nil
	}
}

// ListTagResourcesInvoker 列出绑定到指定资源类型的标签
func (c *OrganizationsClient) ListTagResourcesInvoker(request *model.ListTagResourcesRequest) *ListTagResourcesInvoker {
	requestDef := GenReqDefForListTagResources()
	return &ListTagResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagsForResource 列出绑定到指定资源的标签
//
// 列出绑定到指定资源的标签。您可以将标签附加到组织中的账号、组织单元、根和策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListTagsForResource(request *model.ListTagsForResourceRequest) (*model.ListTagsForResourceResponse, error) {
	requestDef := GenReqDefForListTagsForResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsForResourceResponse), nil
	}
}

// ListTagsForResourceInvoker 列出绑定到指定资源的标签
func (c *OrganizationsClient) ListTagsForResourceInvoker(request *model.ListTagsForResourceRequest) *ListTagsForResourceInvoker {
	requestDef := GenReqDefForListTagsForResource()
	return &ListTagsForResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceInstancesCount 根据资源类型及标签信息查询实例数量
//
// 根据资源类型及标签信息查询实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ShowResourceInstancesCount(request *model.ShowResourceInstancesCountRequest) (*model.ShowResourceInstancesCountResponse, error) {
	requestDef := GenReqDefForShowResourceInstancesCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceInstancesCountResponse), nil
	}
}

// ShowResourceInstancesCountInvoker 根据资源类型及标签信息查询实例数量
func (c *OrganizationsClient) ShowResourceInstancesCountInvoker(request *model.ShowResourceInstancesCountRequest) *ShowResourceInstancesCountInvoker {
	requestDef := GenReqDefForShowResourceInstancesCount()
	return &ShowResourceInstancesCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TagResource 为指定资源添加标签
//
// 向指定的资源添加一个或多个标签。目前，您可以将标签附加到组织中的账号、组织单元、根和策略。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) TagResource(request *model.TagResourceRequest) (*model.TagResourceResponse, error) {
	requestDef := GenReqDefForTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagResourceResponse), nil
	}
}

// TagResourceInvoker 为指定资源添加标签
func (c *OrganizationsClient) TagResourceInvoker(request *model.TagResourceRequest) *TagResourceInvoker {
	requestDef := GenReqDefForTagResource()
	return &TagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UntagResource 从指定资源中删除指定主键标签
//
// 从指定资源中删除具有指定主键的任何标签。您可以将标签绑定到组织中的账号、组织单元、根和策略。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) UntagResource(request *model.UntagResourceRequest) (*model.UntagResourceResponse, error) {
	requestDef := GenReqDefForUntagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UntagResourceResponse), nil
	}
}

// UntagResourceInvoker 从指定资源中删除指定主键标签
func (c *OrganizationsClient) UntagResourceInvoker(request *model.UntagResourceRequest) *UntagResourceInvoker {
	requestDef := GenReqDefForUntagResource()
	return &UntagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableTrustedService 禁用受信任服务
//
// 禁用服务（由service_principal指定的服务）与组织的集成。禁用可信服务后，指定服务将不可以在组织中的新账号中创建服务关联委托。这意味着该服务无法代表您对组织中的任何新账号执行操作。在服务完成从组织中的清理之前，服务仍可以在旧账号中执行操作。此接口只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) DisableTrustedService(request *model.DisableTrustedServiceRequest) (*model.DisableTrustedServiceResponse, error) {
	requestDef := GenReqDefForDisableTrustedService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableTrustedServiceResponse), nil
	}
}

// DisableTrustedServiceInvoker 禁用受信任服务
func (c *OrganizationsClient) DisableTrustedServiceInvoker(request *model.DisableTrustedServiceRequest) *DisableTrustedServiceInvoker {
	requestDef := GenReqDefForDisableTrustedService()
	return &DisableTrustedServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableTrustedService 启用可信服务
//
// 启用服务（由service_principal指定的服务）与组织的集成。启用可信服务后，允许指定的可信服务对组织中的所有账号创建服务关联委托。这允许可信服务代表您在组织及其账号中执行操作。此接口只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) EnableTrustedService(request *model.EnableTrustedServiceRequest) (*model.EnableTrustedServiceResponse, error) {
	requestDef := GenReqDefForEnableTrustedService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableTrustedServiceResponse), nil
	}
}

// EnableTrustedServiceInvoker 启用可信服务
func (c *OrganizationsClient) EnableTrustedServiceInvoker(request *model.EnableTrustedServiceRequest) *EnableTrustedServiceInvoker {
	requestDef := GenReqDefForEnableTrustedService()
	return &EnableTrustedServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrustedServices 列出组织的可信服务列表
//
// 返回启用与组织集成的可信服务列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OrganizationsClient) ListTrustedServices(request *model.ListTrustedServicesRequest) (*model.ListTrustedServicesResponse, error) {
	requestDef := GenReqDefForListTrustedServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrustedServicesResponse), nil
	}
}

// ListTrustedServicesInvoker 列出组织的可信服务列表
func (c *OrganizationsClient) ListTrustedServicesInvoker(request *model.ListTrustedServicesRequest) *ListTrustedServicesInvoker {
	requestDef := GenReqDefForListTrustedServices()
	return &ListTrustedServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
