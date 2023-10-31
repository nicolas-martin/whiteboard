package main

func main() {

}

type BlockchainNetworks []Network

type NovaConfig struct {
	ID                          string
	Endpoint                    string
	SigningStageCreatorApiKeys  []string
	ApprovalStageCreatorApiKeys []string
	EnabledNetworks             BlockchainNetworks
}
type Network int32

const (
	Network_NETWORK_UNKNOWN        Network = 0
	Network_NETWORK_RIPPLE_MAINNET Network = 1
	Network_NETWORK_RIPPLE_TESTNET Network = 2
)
