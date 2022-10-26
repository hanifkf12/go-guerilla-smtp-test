package main

import (
	"fmt"
	"github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg := &guerrilla.AppConfig{
		BackendConfig: backends.BackendConfig{
			"save_process":         "HeadersParser|Debugger|Hasher|Header|Compressor",
			"save_workers_size":    3,
			"sql_driver":           "mysql",
			"sql_dsn":              "root:hanifkf@tcp(127.0.0.1:3306)/goguerrilla?readTimeout=10s&writeTimeout=10s",
			"mail_table":           "new",
			"log_received_mails":   true,
			"primary_mail_host":    "test.org",
			"sql_max_open_conns":   10,
			"sql_max_idle_conns":   5,
			"idle connection pool": 3,
		},
	}
	sc := guerrilla.ServerConfig{
		Hostname:        "test.org",
		ListenInterface: "127.0.0.1:2525",
		IsEnabled:       true,
		MaxClients:      1000,
	}
	cfg.AllowedHosts = []string{"test.org", "gmail.com", "localdomain", "guerrillamail.info", "hanifkf.com"}
	cfg.Servers = append(cfg.Servers, sc)
	d := guerrilla.Daemon{Config: cfg}
	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	}
	d.Log()
	go forever()
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	fmt.Println("Adios!")
}

func forever() {
	for {
		time.Sleep(time.Second)
	}
}
