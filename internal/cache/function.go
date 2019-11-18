package cache

import (
	"fmt"
	"strings"
	"sync"

	"github.com/xymodule/formula/opt"
)

var defaultFunctionCache = make(map[string]*opt.Function, 512)
var rwLock sync.RWMutex

func Register(f *opt.Function) error {
	if f == nil {
		return fmt.Errorf("argument Function can not be nil")
	}

	lowerName := (*f).Name()
	lowerName = strings.ToLower(strings.TrimSpace(lowerName))

	rwLock.Lock()
	defer rwLock.Unlock()

	if _, ok := defaultFunctionCache[lowerName]; ok {
		return fmt.Errorf("function %s duplicate", lowerName)
	}

	defaultFunctionCache[lowerName] = f

	return nil
}

func FindFunction(name string) (*opt.Function, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	if f, ok := defaultFunctionCache[name]; ok {
		return f, nil
	}

	return nil, fmt.Errorf("function %s dot not exists", name)
}
