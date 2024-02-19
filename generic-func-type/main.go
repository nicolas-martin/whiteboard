package main

type (
	Endpoint struct {
		// Define your Endpoint structure here
	}

	CosmosSDK interface {
		// Define the methods or fields expected in CosmosSDK interface
	}

	GenericClientFactoryFn[T any] func(*Endpoint) T
)

func CosmosSDKClientFactory(endpoint *Endpoint) CosmosSDK {
	// Implement the factory logic for CosmosSDK
	// For now, let's assume you have a concrete type that implements the CosmosSDK interface
	return nil
}

func test[T any]() GenericClientFactoryFn[T] {
	return func(endpoint *Endpoint) T {
		return CosmosSDKClientFactory(endpoint).(T)
	}
}

func main() {
	// You can test it like this:
	endpoint := &Endpoint{}
	clientFactory := test[CosmosSDK]()
	client := clientFactory(endpoint)
	// Use the client as needed
	_ = client
}
