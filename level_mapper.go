package logrouter

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type (
	//LevelMapper is a Mapper with key bound to level.Key()
	LevelMapper struct {
		mapper *Mapper
	}
)

func NewLevelMapper() *LevelMapper {
	return &LevelMapper{
		mapper: NewMapper(level.Key()),
	}
}

func (lm *LevelMapper) Log(keyvals ...interface{}) error {
	return lm.mapper.Log(keyvals...)
}

func (lm *LevelMapper) AddDebug(loggers ...log.Logger) *LevelMapper {
	lm.mapper.AddLogger(level.DebugValue(), loggers...)
	return lm
}

func (lm *LevelMapper) AddInfo(loggers ...log.Logger) *LevelMapper {
	lm.mapper.AddLogger(level.InfoValue(), loggers...)
	return lm
}

func (lm *LevelMapper) AddWarn(loggers ...log.Logger) *LevelMapper {
	lm.mapper.AddLogger(level.WarnValue(), loggers...)
	return lm
}

func (lm *LevelMapper) AddError(loggers ...log.Logger) *LevelMapper {
	lm.mapper.AddLogger(level.ErrorValue(), loggers...)
	return lm
}
