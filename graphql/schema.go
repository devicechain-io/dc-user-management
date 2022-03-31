/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	_ "embed"
)

//go:embed schema.gql
var SchemaContent string

type SchemaResolver struct{}
