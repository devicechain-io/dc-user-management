/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package keycloak

import (
	"context"
	"fmt"

	gocloak "github.com/Nerzal/gocloak/v11"
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/rs/zerolog/log"
)

// Manages lifecycle of Keycloak interactions.
type KeycloakManager struct {
	Microservice *core.Microservice

	// Keycloak client.
	Keycloak gocloak.GoCloak
	JWT      *gocloak.JWT

	lifecycle core.LifecycleManager
}

// Create a new Keycloak manager.
func NewKeycloakManager(ms *core.Microservice, callbacks core.LifecycleCallbacks) *KeycloakManager {
	kc := &KeycloakManager{
		Microservice: ms,
	}

	// Create lifecycle manager.
	name := fmt.Sprintf("%s-%s", ms.FunctionalArea, "keycloak")
	kc.lifecycle = core.NewLifecycleManager(name, kc, callbacks)
	return kc
}

// Initialize component.
func (kmgr *KeycloakManager) Initialize(ctx context.Context) error {
	return kmgr.lifecycle.Initialize(ctx)
}

// Lifecycle callback that runs initialization logic.
func (kmgr *KeycloakManager) ExecuteInitialize(ctx context.Context) error {
	kconfig := kmgr.Microservice.InstanceConfiguration.Infrastructure.Keycloak
	url := fmt.Sprintf("%s://%s:%d", "http", kconfig.Hostname, 8080)

	kmgr.Keycloak = gocloak.NewClient(url)
	jwt, err := kmgr.Keycloak.LoginAdmin(ctx, "devicechain", "devicechain", "master")
	if err != nil {
		return err
	}
	kmgr.JWT = jwt
	log.Info().Msg("Logged in to Keycloak master realm successfully.")

	return nil
}

// Start component.
func (kmgr *KeycloakManager) Start(ctx context.Context) error {
	return kmgr.lifecycle.Start(ctx)
}

// Lifecycle callback that runs startup logic.
func (kmgr *KeycloakManager) ExecuteStart(context.Context) error {
	return nil
}

// Stop component.
func (kmgr *KeycloakManager) Stop(ctx context.Context) error {
	return kmgr.lifecycle.Stop(ctx)
}

// Lifecycle callback that runs shutdown logic.
func (kmgr *KeycloakManager) ExecuteStop(context.Context) error {
	return nil
}

// Terminate component.
func (kmgr *KeycloakManager) Terminate(ctx context.Context) error {
	return kmgr.lifecycle.Terminate(ctx)
}

// Lifecycle callback that runs termination logic.
func (kmgr *KeycloakManager) ExecuteTerminate(context.Context) error {
	return nil
}
