/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package generator

import (
	"fmt"

	gen "github.com/devicechain-io/dc-k8s/generators"
	"github.com/devicechain-io/dc-user-management/config"
)

type ResourceProvider struct{}

// Get instance configuration CRs that should be created in tooling
func (ResourceProvider) GetConfigurationResources() ([]gen.ConfigurationResource, error) {
	resources := make([]gen.ConfigurationResource, 0)

	name := "user-management-default"
	msconfig := config.NewUserManagementConfiguration()
	content, err := gen.GenerateMicroserviceConfig(name, "user-management", "devicechain-io/user-management:0.0.1", msconfig)
	if err != nil {
		return nil, err
	}
	dcmdefault := gen.ConfigurationResource{
		Name:    fmt.Sprintf("%s_%s", "core.devicechain.io", name),
		Content: content,
	}

	resources = append(resources, dcmdefault)
	return resources, nil
}
