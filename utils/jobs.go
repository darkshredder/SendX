package utils

import "fmt"

type Job struct {
	Name      string
	Retries   int
	Functions func() bool
	Tries     int
}

func DoWork(id int, j Job, jobs chan Job) {
	fmt.Printf("worker%d: started %s", id, j.Name)
	resp := j.Functions()
	if resp == false {
		fmt.Printf("worker%d: error %s\n", id, j.Name)
		if j.Tries < j.Retries {
			j.Tries++
			jobs <- j
		}
	}
	fmt.Printf("worker%d: completed %s!\n", id, j.Name)
	return

}
