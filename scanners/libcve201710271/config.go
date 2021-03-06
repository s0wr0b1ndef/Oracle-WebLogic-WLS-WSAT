package libcve201710271

// Config is used to store the different configurations that we need
type Config struct {
	Lhost      string
	Lport      int
	TargetFile string
	Verbose    bool
	OutputFile string
	Threads    int
	WaitTime   int
	AllURLs    bool
}
