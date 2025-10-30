package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKvsCredentialsBuilder_SafeBuild(t *testing.T) {
	credentials, err := NewKvsCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("token").
		WithIdTokenFile("file").
		WithIdpId("idp-id").
		WithIamEndpointOverride("iam-endpoint").
		WithProjectId("project-id").
		SafeBuild()
	assert.NoError(t, err)
	assert.IsType(t, &KvsCredentials{}, credentials)
	assert.Equal(t, "ak", credentials.AK)
	assert.Equal(t, "sk", credentials.SK)
	assert.Equal(t, "token", credentials.SecurityToken)
	assert.Equal(t, "idp-id", credentials.IdpId)
	assert.Equal(t, "project-id", credentials.ProjectId)
	assert.Equal(t, "iam-endpoint", credentials.IamEndpoint)
	assert.Equal(t, "file", credentials.IdTokenFile)

	credentials.AK = "new_ak"
	credentials.ProjectId = "new-project-id"
	assert.Equal(t, "new_ak", credentials.AK)
	assert.Equal(t, "new-project-id", credentials.ProjectId)
}
