/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package main

import (
	"context"

	gql "github.com/graph-gophers/graphql-go"

	"github.com/devicechain-io/dc-microservice/core"
	gqlcore "github.com/devicechain-io/dc-microservice/graphql"
	"github.com/devicechain-io/dc-user-management/graphql"
	"github.com/devicechain-io/dc-user-management/keycloak"
)

var (
	Microservice *core.Microservice

	GraphQLManager  *gqlcore.GraphQLManager
	KeycloakManager *keycloak.KeycloakManager
)

func main() {
	callbacks := core.LifecycleCallbacks{
		Initializer: core.LifecycleCallback{
			Preprocess:  func(context.Context) error { return nil },
			Postprocess: afterMicroserviceInitialized,
		},
		Starter: core.LifecycleCallback{
			Preprocess:  func(context.Context) error { return nil },
			Postprocess: afterMicroserviceStarted,
		},
		Stopper: core.LifecycleCallback{
			Preprocess:  beforeMicroserviceStopped,
			Postprocess: func(context.Context) error { return nil },
		},
		Terminator: core.LifecycleCallback{
			Preprocess:  beforeMicroserviceTerminated,
			Postprocess: func(context.Context) error { return nil },
		},
	}
	Microservice = core.NewMicroservice(callbacks)
	Microservice.Run()
}

// Called after microservice has been initialized.
func afterMicroserviceInitialized(ctx context.Context) error {
	// Create and initialize Keycloak manager.
	KeycloakManager = keycloak.NewKeycloakManager(Microservice, core.NewNoOpLifecycleCallbacks())
	err := KeycloakManager.Initialize(ctx)
	if err != nil {
		return err
	}

	// Map of providers that will be injected into graphql http context.
	providers := map[gqlcore.ContextKey]interface{}{}

	// Create and initialize graphql manager.
	schema := gqlcore.CommonTypes + graphql.SchemaContent
	parsed := gql.MustParseSchema(schema, &graphql.SchemaResolver{})
	GraphQLManager = gqlcore.NewGraphQLManager(Microservice, core.NewNoOpLifecycleCallbacks(), *parsed, providers)
	err = GraphQLManager.Initialize(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Called after microservice has been started.
func afterMicroserviceStarted(ctx context.Context) error {
	// Start Keycloak manager.
	err := KeycloakManager.Start(ctx)
	if err != nil {
		return err
	}

	// Start graphql manager.
	err = GraphQLManager.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Called before microservice has been stopped.
func beforeMicroserviceStopped(ctx context.Context) error {
	// Stop graphql manager.
	err := GraphQLManager.Stop(ctx)
	if err != nil {
		return err
	}

	// Stop Keycloak manager.
	err = KeycloakManager.Stop(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Called before microservice has been terminated.
func beforeMicroserviceTerminated(ctx context.Context) error {
	// Terminate graphql manager.
	err := GraphQLManager.Terminate(ctx)
	if err != nil {
		return err
	}

	// Terminate Keycloak manager.
	err = KeycloakManager.Terminate(ctx)
	if err != nil {
		return err
	}

	return nil
}
