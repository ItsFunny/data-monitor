package constants

const (
	MAC_ADDR_LEN = 6
	ETH_TYPE_LEN = 2
)

const (
	// other
	MAX_TCP_OPTION_SIZE = 40

	// header size
	ETHERNET_HEADER_SIZE = MAC_ADDR_LEN*2 + ETH_TYPE_LEN
	ARP_HEADER_SIZE      = 28 //
	IP_HEADER_SIZE       = 20 // 20
)

// offset
const (
	OFFSET_DST_ADD  = 0
	OFFSET_SRC_ADD  = OFFSET_DST_ADD + MAC_ADDR_LEN
	OFFSET_ETH_TYPE = OFFSET_SRC_ADD + MAC_ADDR_LEN
	// 版头服包
	OFFSET_IP_PROTOCOL = 23 //0x800 0x86
	OFFSET_SRC_IP      = 26 //
	OFFSET_DST_IP      = 30
	OFFSET_DST_PORT    = 36 // 14+20+2

)
