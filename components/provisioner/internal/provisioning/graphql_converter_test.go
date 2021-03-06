package provisioning

import (
	"testing"

	"github.com/kyma-project/control-plane/components/provisioner/internal/util"

	"github.com/stretchr/testify/require"

	"github.com/kyma-project/control-plane/components/provisioner/internal/model"
	"github.com/kyma-project/control-plane/components/provisioner/pkg/gqlschema"
	"github.com/stretchr/testify/assert"
)

const (
	kymaSystemNamespace      = "kyma-system"
	kymaIntegrationNamespace = "kyma-integration"
)

func TestOperationStatusToGQLOperationStatus(t *testing.T) {

	graphQLConverter := NewGraphQLConverter()

	t.Run("Should create proper operation status struct", func(t *testing.T) {
		//given
		operation := model.Operation{
			ID:        "5f6e3ab6-d803-430a-8fac-29c9c9b4485a",
			Type:      model.Upgrade,
			State:     model.InProgress,
			Message:   "Some message",
			ClusterID: "6af76034-272a-42be-ac39-30e075f515a3",
		}

		operationID := "5f6e3ab6-d803-430a-8fac-29c9c9b4485a"
		message := "Some message"
		runtimeID := "6af76034-272a-42be-ac39-30e075f515a3"

		expectedOperationStatus := &gqlschema.OperationStatus{
			ID:        &operationID,
			Operation: gqlschema.OperationTypeUpgrade,
			State:     gqlschema.OperationStateInProgress,
			Message:   &message,
			RuntimeID: &runtimeID,
		}

		//when
		status := graphQLConverter.OperationStatusToGQLOperationStatus(operation)

		//then
		assert.Equal(t, expectedOperationStatus, status)
	})
}

func TestRuntimeStatusToGraphQLStatus(t *testing.T) {

	graphQLConverter := NewGraphQLConverter()

	t.Run("Should create proper runtime status struct for gardener config with zones", func(t *testing.T) {
		//given
		clusterName := "Something"
		project := "Project"
		disk := "standard"
		machine := "machine"
		region := "region"
		zones := []string{"fix-gcp-zone-1", "fix-gcp-zone-2"}
		volume := 256
		kubeversion := "kubeversion"
		kubeconfig := "kubeconfig"
		provider := "GCP"
		purpose := "testing"
		licenceType := "partner"
		seed := "gcp-eu1"
		secret := "secret"
		cidr := "cidr"
		autoScMax := 2
		autoScMin := 2
		surge := 1
		unavailable := 1

		gardenerProviderConfig, err := model.NewGardenerProviderConfigFromJSON(`{"zones":["fix-gcp-zone-1","fix-gcp-zone-2"]}`)
		require.NoError(t, err)

		runtimeStatus := model.RuntimeStatus{
			LastOperationStatus: model.Operation{
				ID:        "5f6e3ab6-d803-430a-8fac-29c9c9b4485a",
				Type:      model.Deprovision,
				State:     model.Failed,
				Message:   "Some message",
				ClusterID: "6af76034-272a-42be-ac39-30e075f515a3",
			},
			RuntimeConnectionStatus: model.RuntimeAgentConnectionStatusDisconnected,
			RuntimeConfiguration: model.Cluster{
				ClusterConfig: model.GardenerConfig{
					Name:                   clusterName,
					ProjectName:            project,
					DiskType:               disk,
					MachineType:            machine,
					Region:                 region,
					VolumeSizeGB:           volume,
					KubernetesVersion:      kubeversion,
					Provider:               provider,
					Purpose:                &purpose,
					LicenceType:            &licenceType,
					Seed:                   seed,
					TargetSecret:           secret,
					WorkerCidr:             cidr,
					AutoScalerMax:          autoScMax,
					AutoScalerMin:          autoScMin,
					MaxSurge:               surge,
					MaxUnavailable:         unavailable,
					GardenerProviderConfig: gardenerProviderConfig,
				},
				Kubeconfig: &kubeconfig,
				KymaConfig: fixKymaConfig(),
			},
		}

		operationID := "5f6e3ab6-d803-430a-8fac-29c9c9b4485a"
		message := "Some message"
		runtimeID := "6af76034-272a-42be-ac39-30e075f515a3"

		expectedRuntimeStatus := &gqlschema.RuntimeStatus{
			LastOperationStatus: &gqlschema.OperationStatus{
				ID:        &operationID,
				Operation: gqlschema.OperationTypeDeprovision,
				State:     gqlschema.OperationStateFailed,
				Message:   &message,
				RuntimeID: &runtimeID,
			},
			RuntimeConnectionStatus: &gqlschema.RuntimeConnectionStatus{
				Status: gqlschema.RuntimeAgentConnectionStatusDisconnected,
			},
			RuntimeConfiguration: &gqlschema.RuntimeConfig{
				ClusterConfig: &gqlschema.GardenerConfig{
					Name:              &clusterName,
					DiskType:          &disk,
					MachineType:       &machine,
					Region:            &region,
					VolumeSizeGb:      &volume,
					KubernetesVersion: &kubeversion,
					Provider:          &provider,
					Purpose:           &purpose,
					LicenceType:       &licenceType,
					Seed:              &seed,
					TargetSecret:      &secret,
					WorkerCidr:        &cidr,
					AutoScalerMax:     &autoScMax,
					AutoScalerMin:     &autoScMin,
					MaxSurge:          &surge,
					MaxUnavailable:    &unavailable,
					ProviderSpecificConfig: gqlschema.GCPProviderConfig{
						Zones: zones,
					},
				},
				KymaConfig: fixKymaGraphQLConfig(),
				Kubeconfig: &kubeconfig,
			},
		}

		//when
		gqlStatus := graphQLConverter.RuntimeStatusToGraphQLStatus(runtimeStatus)

		//then
		assert.Equal(t, expectedRuntimeStatus, gqlStatus)
	})

	t.Run("Should create proper runtime status struct for gardener config without zones", func(t *testing.T) {
		//given
		clusterName := "Something"
		project := "Project"
		disk := "standard"
		machine := "machine"
		region := "region"
		volume := 256
		kubeversion := "kubeversion"
		kubeconfig := "kubeconfig"
		provider := "Azure"
		purpose := "testing"
		licenceType := ""
		seed := "az-eu1"
		secret := "secret"
		cidr := "cidr"
		autoScMax := 2
		autoScMin := 2
		surge := 1
		unavailable := 1

		gardenerProviderConfig, err := model.NewGardenerProviderConfigFromJSON(`{"vnetCidr":"10.10.11.11/255"}`)
		require.NoError(t, err)

		runtimeStatus := model.RuntimeStatus{
			LastOperationStatus: model.Operation{
				ID:        "5f6e3ab6-d803-430a-8fac-29c9c9b4485a",
				Type:      model.Deprovision,
				State:     model.Failed,
				Message:   "Some message",
				ClusterID: "6af76034-272a-42be-ac39-30e075f515a3",
			},
			RuntimeConnectionStatus: model.RuntimeAgentConnectionStatusDisconnected,
			RuntimeConfiguration: model.Cluster{
				ClusterConfig: model.GardenerConfig{
					Name:                   clusterName,
					ProjectName:            project,
					KubernetesVersion:      kubeversion,
					VolumeSizeGB:           volume,
					DiskType:               disk,
					MachineType:            machine,
					Provider:               provider,
					Purpose:                &purpose,
					LicenceType:            &licenceType,
					Seed:                   seed,
					TargetSecret:           secret,
					Region:                 region,
					WorkerCidr:             cidr,
					AutoScalerMin:          autoScMin,
					AutoScalerMax:          autoScMax,
					MaxSurge:               surge,
					MaxUnavailable:         unavailable,
					GardenerProviderConfig: gardenerProviderConfig,
				},
				Kubeconfig: &kubeconfig,
				KymaConfig: fixKymaConfig(),
			},
		}

		operationID := "5f6e3ab6-d803-430a-8fac-29c9c9b4485a"
		message := "Some message"
		runtimeID := "6af76034-272a-42be-ac39-30e075f515a3"

		expectedRuntimeStatus := &gqlschema.RuntimeStatus{
			LastOperationStatus: &gqlschema.OperationStatus{
				ID:        &operationID,
				Operation: gqlschema.OperationTypeDeprovision,
				State:     gqlschema.OperationStateFailed,
				Message:   &message,
				RuntimeID: &runtimeID,
			},
			RuntimeConnectionStatus: &gqlschema.RuntimeConnectionStatus{
				Status: gqlschema.RuntimeAgentConnectionStatusDisconnected,
			},
			RuntimeConfiguration: &gqlschema.RuntimeConfig{
				ClusterConfig: &gqlschema.GardenerConfig{
					Name:              &clusterName,
					DiskType:          &disk,
					MachineType:       &machine,
					Region:            &region,
					VolumeSizeGb:      &volume,
					KubernetesVersion: &kubeversion,
					Provider:          &provider,
					Purpose:           &purpose,
					LicenceType:       &licenceType,
					Seed:              &seed,
					TargetSecret:      &secret,
					WorkerCidr:        &cidr,
					AutoScalerMax:     &autoScMax,
					AutoScalerMin:     &autoScMin,
					MaxSurge:          &surge,
					MaxUnavailable:    &unavailable,
					ProviderSpecificConfig: gqlschema.AzureProviderConfig{
						VnetCidr: util.StringPtr("10.10.11.11/255"),
						Zones:    nil, // Expected empty when no zones specified in input.
					},
				},
				KymaConfig: fixKymaGraphQLConfig(),
				Kubeconfig: &kubeconfig,
			},
		}

		//when
		gqlStatus := graphQLConverter.RuntimeStatusToGraphQLStatus(runtimeStatus)

		//then
		assert.Equal(t, expectedRuntimeStatus, gqlStatus)
	})
}

func fixKymaGraphQLConfig() *gqlschema.KymaConfig {
	return &gqlschema.KymaConfig{
		Version: util.StringPtr(kymaVersion),
		Components: []*gqlschema.ComponentConfiguration{
			{
				Component:     clusterEssentialsComponent,
				Namespace:     kymaSystemNamespace,
				Configuration: make([]*gqlschema.ConfigEntry, 0, 0),
			},
			{
				Component: coreComponent,
				Namespace: kymaSystemNamespace,
				Configuration: []*gqlschema.ConfigEntry{
					fixGQLConfigEntry("test.config.key", "value", util.BoolPtr(false)),
					fixGQLConfigEntry("test.config.key2", "value2", util.BoolPtr(false)),
				},
			},
			{
				Component:     rafterComponent,
				Namespace:     kymaSystemNamespace,
				SourceURL:     util.StringPtr(rafterSourceURL),
				Configuration: make([]*gqlschema.ConfigEntry, 0, 0),
			},
			{
				Component: applicationConnectorComponent,
				Namespace: kymaIntegrationNamespace,
				Configuration: []*gqlschema.ConfigEntry{
					fixGQLConfigEntry("test.config.key", "value", util.BoolPtr(false)),
					fixGQLConfigEntry("test.secret.key", "secretValue", util.BoolPtr(true)),
				},
			},
		},
		Configuration: []*gqlschema.ConfigEntry{
			fixGQLConfigEntry("global.config.key", "globalValue", util.BoolPtr(false)),
			fixGQLConfigEntry("global.config.key2", "globalValue2", util.BoolPtr(false)),
			fixGQLConfigEntry("global.secret.key", "globalSecretValue", util.BoolPtr(true)),
		},
	}
}

func fixGQLConfigEntry(key, val string, secret *bool) *gqlschema.ConfigEntry {
	return &gqlschema.ConfigEntry{
		Key:    key,
		Value:  val,
		Secret: secret,
	}
}

func fixKymaConfig() model.KymaConfig {
	return model.KymaConfig{
		ID:                  "id",
		Release:             fixKymaRelease(),
		Components:          fixKymaComponents(),
		GlobalConfiguration: fixGlobalConfig(),
		ClusterID:           "runtimeID",
	}
}

func fixGlobalConfig() model.Configuration {
	return model.Configuration{
		ConfigEntries: []model.ConfigEntry{
			model.NewConfigEntry("global.config.key", "globalValue", false),
			model.NewConfigEntry("global.config.key2", "globalValue2", false),
			model.NewConfigEntry("global.secret.key", "globalSecretValue", true),
		},
	}
}

func fixKymaComponents() []model.KymaComponentConfig {
	return []model.KymaComponentConfig{
		{
			ID:             "id",
			KymaConfigID:   "id",
			Component:      clusterEssentialsComponent,
			Namespace:      kymaSystemNamespace,
			Configuration:  model.Configuration{ConfigEntries: make([]model.ConfigEntry, 0, 0)},
			ComponentOrder: 1,
		},
		{
			ID:           "id",
			KymaConfigID: "id",
			Component:    coreComponent,
			Namespace:    kymaSystemNamespace,
			Configuration: model.Configuration{
				ConfigEntries: []model.ConfigEntry{
					model.NewConfigEntry("test.config.key", "value", false),
					model.NewConfigEntry("test.config.key2", "value2", false),
				},
			},
			ComponentOrder: 2,
		},
		{
			ID:             "id",
			KymaConfigID:   "id",
			Component:      rafterComponent,
			Namespace:      kymaSystemNamespace,
			SourceURL:      util.StringPtr(rafterSourceURL),
			Configuration:  model.Configuration{ConfigEntries: make([]model.ConfigEntry, 0, 0)},
			ComponentOrder: 3,
		},
		{
			ID:           "id",
			KymaConfigID: "id",
			Component:    applicationConnectorComponent,
			Namespace:    kymaIntegrationNamespace,
			Configuration: model.Configuration{
				ConfigEntries: []model.ConfigEntry{
					model.NewConfigEntry("test.config.key", "value", false),
					model.NewConfigEntry("test.secret.key", "secretValue", true),
				},
			},
			ComponentOrder: 4,
		},
	}
}

func fixKymaRelease() model.Release {
	return model.Release{
		Id:            "d829b1b5-2e82-426d-91b0-f94978c0c140",
		Version:       kymaVersion,
		TillerYAML:    "tiller yaml",
		InstallerYAML: "installer yaml",
	}
}
