package CastHunter

type Flags struct {
	Arp                 bool
	SendOutEvery        int
	OneInterface        bool
	Server              bool   // launch a local HTTP server
	Trace               bool   // Traceback logging for panic
	Error               bool   // Error logging for panic
	SSDP                bool   // Enable SSDP on a network card
	NetworkInterface    string // For SSDP listeners
	MetworkInterfaceARP string // disable auto listen for cards and choose a manual interface for ARP
	Help                bool   // Help menu
	SuperHelp           bool   // More descriptive help menu
}
