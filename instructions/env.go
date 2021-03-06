package instructions

import (
	"github.com/thehivecorporation/raccoon/connection"

	log "github.com/Sirupsen/logrus"

	"strings"
)

type ENV struct {
	Name        string
	Description string
	Environment string
}

func (e *ENV) Execute(n connection.Node) {
	session, err := n.GetSession()
	if err != nil {
		log.WithFields(log.Fields{
			"Instruction": "RUN",
			"Node":        n.IP,
			"package":     "instructions",
		}).Error(err.Error())

		session.Close()
		return
	}

	log.WithFields(log.Fields{
		"Instruction": "ENV",
		"Node":        n.IP,
		"package":     "instructions",
	}).Info(e.Description)

	env := strings.Split(e.Environment, "=")
	if len(env) == 2 {
		session.Setenv(env[0], env[1])
	} else {
		log.Fatal("More than one '=' found on ENV instruction")
	}
}
