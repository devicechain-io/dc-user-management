/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package config

type NestedConfiguration struct {
	Test string
}

type UserManagementConfiguration struct {
	Nested NestedConfiguration
}

// Creates the default device management configuration
func NewUserManagementConfiguration() *UserManagementConfiguration {
	return &UserManagementConfiguration{
		Nested: NestedConfiguration{
			Test: "test",
		},
	}
}
