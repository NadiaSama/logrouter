package logrouter

import (
	"errors"

	"github.com/go-kit/log"
)

type (
	//Mapper route log event according to the value of key in log keyvals
	//if no associate logger was found the def logger will be used or ErrNoLogger
	//is return
	Mapper struct {
		valMap map[interface{}][]log.Logger
		def    log.Logger
		key    interface{}
	}
)

var (
	//ErrNoLogger means no associate logger for keyvals
	ErrNoLogger = errors.New("no available logger")
)

//New return a Mapper with specific key
func NewMapper(key interface{}) *Mapper {
	return &Mapper{
		valMap: make(map[interface{}][]log.Logger),
		key:    key,
	}
}

func (m *Mapper) SetDefault(def log.Logger) *Mapper {
	m.def = def
	return m
}

func (m *Mapper) AddLogger(val interface{}, loggers ...log.Logger) *Mapper {
	ls, ok := m.valMap[val]
	if !ok {
		ls = []log.Logger{}
		m.valMap[val] = ls
	}

	ls = append(ls, loggers...)
	m.valMap[val] = ls
	return m
}

func (m *Mapper) Log(keyvals ...interface{}) error {
	l := len(keyvals) - 1 //avoid odd condition
	for i := 0; i < l; i += 2 {
		k := keyvals[i]

		if k != m.key {
			continue
		}

		v := keyvals[i+1]

		ls, ok := m.valMap[v]
		if ok {
			for _, l := range ls {
				if err := l.Log(keyvals...); err != nil {
					return err
				}
			}
		}

		if m.def == nil {
			return ErrNoLogger
		}
		return m.def.Log(keyvals...)
	}

	if m.def == nil {
		return ErrNoLogger
	}
	return m.def.Log(keyvals...)
}
