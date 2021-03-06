// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Provider for the core.zephyr.solo.io/v1alpha1 Clientset from config
func ClientsetFromConfigProvider(cfg *rest.Config) (Clientset, error) {
	return NewClientsetFromConfig(cfg)
}

// Provider for the core.zephyr.solo.io/v1alpha1 Clientset from client
func ClientsProvider(client client.Client) Clientset {
	return NewClientset(client)
}

// Provider for SettingsClient from Clientset
func SettingsClientFromClientsetProvider(clients Clientset) SettingsClient {
	return clients.Settings()
}

// Provider for SettingsClient from Client
func SettingsClientProvider(client client.Client) SettingsClient {
	return NewSettingsClient(client)
}

type SettingsClientFactory func(client client.Client) SettingsClient

func SettingsClientFactoryProvider() SettingsClientFactory {
	return SettingsClientProvider
}

type SettingsClientFromConfigFactory func(cfg *rest.Config) (SettingsClient, error)

func SettingsClientFromConfigFactoryProvider() SettingsClientFromConfigFactory {
	return func(cfg *rest.Config) (SettingsClient, error) {
		clients, err := NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Settings(), nil
	}
}
