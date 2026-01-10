package common

type NetworkConfig struct {
	Name    string
	ChainID int64

	GatewayREST     string
	GatewayV2       string
	GatewayWS       string
	EndPoint        string
	SubscriptionsWS string

	ArchiveREST string
	ArchiveV2   string

	TriggerREST string
}

const (
	PathExecute = "/execute"
	PathQuery   = "/query"
)

var (
	Mainnet = NetworkConfig{
		Name:    "Mainnet",
		ChainID: 57073,

		GatewayREST: "https://gateway.prod.nado.xyz/v1",
		GatewayV2:   "https://gateway.prod.nado.xyz/v2",
		GatewayWS:   "wss://gateway.prod.nado.xyz/v1/ws",

		SubscriptionsWS: "wss://gateway.prod.nado.xyz/v1/subscribe",

		ArchiveREST: "https://archive.prod.nado.xyz/v1",
		ArchiveV2:   "https://archive.prod.nado.xyz/v2",

		TriggerREST: "https://trigger.prod.nado.xyz/v1",
		EndPoint:    "0x05ec92d78ed421f3d3ada77ffde167106565974e",
	}

	Testnet = NetworkConfig{
		Name:    "Testnet (Ink Sepolia)",
		ChainID: 763373,

		GatewayREST: "https://gateway.test.nado.xyz/v1",
		GatewayV2:   "https://gateway.test.nado.xyz/v2",
		GatewayWS:   "wss://gateway.test.nado.xyz/v1/ws",

		SubscriptionsWS: "wss://gateway.test.nado.xyz/v1/subscribe",

		ArchiveREST: "https://archive.test.nado.xyz/v1",
		ArchiveV2:   "https://archive.test.nado.xyz/v2",

		TriggerREST: "https://trigger.test.nado.xyz/v1",
	}
)
