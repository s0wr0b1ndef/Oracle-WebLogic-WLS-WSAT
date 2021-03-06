package libcve201710271

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// Worker is used to create a new worker which we'll use when sending requests
func Worker(id int, m *sync.Mutex, jobs <-chan TargetHost) {
	m.Lock()
	logrus.Infof("Worker %d started", id)
	m.Unlock()
	for th := range jobs {
		SendRequest(th, id, m)
	}
	m.Lock()
	logrus.Infof("Worker %d finished", id)
	m.Unlock()
}
