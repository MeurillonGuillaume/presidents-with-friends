package wsserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Start()
}

type PresidentConfig struct {
	Port int `required:"true"`
}

type PresidentServer struct {
	cfg      PresidentConfig
	upgrader websocket.Upgrader
}

type PresidentMessage struct {
	Message string `json:"message"`
}

func NewPresidentServer(cfg PresidentConfig) Server {
	return &PresidentServer{
		cfg:      cfg,
		upgrader: websocket.Upgrader{},
	}
}

func (ps *PresidentServer) Start() {
	http.HandleFunc("/", ps.status)
	http.HandleFunc("/presidents", ps.listen)
	logrus.Info("Server is online")
	log.Fatalf("Could not continue server loop: %v", http.ListenAndServe(fmt.Sprintf(":%d", ps.cfg.Port), nil))
}

func (ps *PresidentServer) status(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("I'm alive")); err != nil {
		logrus.Errorf("Could not write status response: %v", err)
	}
}

func (ps *PresidentServer) listen(w http.ResponseWriter, r *http.Request) {
	connection, err := ps.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Could not upgrade connection to websockets: %v", err)
	}
	defer connection.Close()

	for {
		var msg PresidentMessage
		if err := connection.ReadJSON(&msg); err != nil {
			logrus.Errorf("Could not read JSON message: %v", err)
		} else {
			logrus.Info(msg)
		}
	}
}
