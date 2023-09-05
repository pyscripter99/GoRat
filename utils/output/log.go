package output

import (
	"go-rat/utils/net"
	"go-rat/utils/types"
	"time"
)

func logTask(task_id uint, log types.Log) error {
	err := net.PostJSON("/log/task/", &log, nil)
	if err != nil {
		return err
	}
	return nil
}

func LogDebug(task_id uint, body string) error {
	return logTask(task_id, types.Log{
		Level: 0,
		Body:  body,
		Time:  time.Now(),
	})
}

func LogInfo(task_id uint, body string) error {
	return logTask(task_id, types.Log{
		Level: 1,
		Body:  body,
		Time:  time.Now(),
	})
}

func LogWarn(task_id uint, body string) error {
	return logTask(task_id, types.Log{
		Level: 2,
		Body:  body,
		Time:  time.Now(),
	})
}
func LogError(task_id uint, body string) error {
	return logTask(task_id, types.Log{
		Level: 3,
		Body:  body,
		Time:  time.Now(),
	})
}

func LogFatal(task_id uint, body string) error {
	return logTask(task_id, types.Log{
		Level: 4,
		Body:  body,
		Time:  time.Now(),
	})
}
