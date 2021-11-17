package msrandom

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/com-gft-tsbo-source/go-common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
// MsRandom
// ###########################################################################
// ###########################################################################

// MsRandom Encapsulates the ms-random data
type MsRandom struct {
	microservice.MicroService

	seededRand  *rand.Rand
	randomMutex sync.Mutex
}

// ###########################################################################

// InitMsRandomFromArgs ...
func InitFromArgs(ms *MsRandom, args []string, flagset *flag.FlagSet) *MsRandom {
	var cfg Configuration

	if flagset == nil {
		flagset = flag.NewFlagSet("ms-random", flag.PanicOnError)
	}

	InitConfigurationFromArgs(&cfg, args, flagset)
	microservice.Init(&ms.MicroService, &cfg.Configuration, nil)
	ms.seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomHandler := ms.DefaultHandler()
	randomHandler.Get = ms.httpGetRandom
	ms.AddHandler("/random", randomHandler)
	return ms
}

// ---------------------------------------------------------------------------

var deviceMutex sync.Mutex

func (ms *MsRandom) httpGetRandom(w http.ResponseWriter, r *http.Request) (status int, contentLen int, msg string) {
	ms.randomMutex.Lock()
	value := ms.seededRand.Intn(100)
	ms.randomMutex.Unlock()
	status = http.StatusOK
	name := r.Header.Get("X-Cid")
	version := r.Header.Get("X-Version")
	environment := r.Header.Get("X-Environment")
	msg = fmt.Sprintf("'v%s' in '%s' Generated random number '%d' for client '%s@%s'.", ms.GetVersion(), environment, value, name, version)
	response := NewRandomResponse(msg, ms)
	response.Value = value
	ms.SetResponseHeaders("application/json; charset=utf-8", w, r)
	w.WriteHeader(status)
	contentLen = ms.Reply(w, response)
	return status, contentLen, msg
}
