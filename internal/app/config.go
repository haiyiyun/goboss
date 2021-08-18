package app

import "github.com/haiyiyun/config"

type Config struct {
	MaxProcs               int             `json:"max_procs"`
	BindAddr               string          `json:"bind_addr"`
	ServerReadTimeOut      config.Duration `json:"server_read_time_out"`
	ServerWriteTimeOut     config.Duration `json:"server_write_time_out"`
	SslCertFile            string          `json:"ssl_cert_file"`
	SslKeyFile             string          `json:"ssl_key_file"`
	LocateRelativeExecPath bool            `json:"locate_relative_exec_path"`
	LogDir                 string          `json:"log_dir"`
	LogFile                string          `json:"log_file"`
	LogDebugLevel          string          `json:"log_debug_level"`
}
