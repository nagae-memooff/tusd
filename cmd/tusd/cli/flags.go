package cli

import (
	"flag"
	"path/filepath"
)

var Flags struct {
	HttpHost      string
	HttpPort      string
	MaxSize       int64
	UploadDir     string
	StoreSize     int64
	Basepath      string
	Timeout       int64
	HooksDir      string
	ShowVersion   bool
	ExposeMetrics bool
	MetricsPath   string
	BehindProxy   bool

	HooksInstalled bool
}

func ParseFlags() {
	flag.StringVar(&Flags.HttpHost, "host", "0.0.0.0", "Host to bind HTTP server to")
	flag.StringVar(&Flags.HttpPort, "port", "8088", "Port to bind HTTP server to")
	flag.Int64Var(&Flags.MaxSize, "max-size", 0, "Maximum size of a single upload in bytes")
	flag.StringVar(&Flags.UploadDir, "dir", "./data", "Directory to store uploads in")
	flag.Int64Var(&Flags.StoreSize, "store-size", 0, "Size of space allowed for storage")
	flag.StringVar(&Flags.Basepath, "base-path", "/files/", "Basepath of the HTTP server")
	flag.Int64Var(&Flags.Timeout, "timeout", 30*1000, "Read timeout for connections in milliseconds.  A zero value means that reads will not timeout")
	flag.StringVar(&Flags.HooksDir, "hooks-dir", "", "Directory to search for available hooks scripts")
	flag.BoolVar(&Flags.ShowVersion, "version", false, "Print tusd version information")
	flag.BoolVar(&Flags.ExposeMetrics, "expose-metrics", true, "Expose metrics about tusd usage")
	flag.StringVar(&Flags.MetricsPath, "metrics-path", "/metrics", "Path under which the metrics endpoint will be accessible")
	flag.BoolVar(&Flags.BehindProxy, "behind-proxy", false, "Respect X-Forwarded-* and similar headers which may be set by proxies")

	flag.Parse()

	if Flags.HooksDir != "" {
		Flags.HooksDir, _ = filepath.Abs(Flags.HooksDir)
		Flags.HooksInstalled = true

		stdout.Printf("Using '%s' for hooks", Flags.HooksDir)
	}

	if Flags.UploadDir == "" {
		stderr.Fatalf("no data dir specified.")
	}
}
