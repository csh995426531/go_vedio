package taskrunner

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestRunner(t *testing.T) {
	t.Run("runner", testRunner)
}

func testRunner(t *testing.T) {
	// fmt.Println("开始\r\n")

	runner := NewRunner(3, true, runnerDispatcher, runnerExecutor)

	runner.StartAll()
}

func runnerDispatcher(dc dataChan) error {

	for i := 0; i < 3; i++ {
		dc <- i
		fmt.Printf("d is: %v \r\n", i)
	}
	return nil
}

func runnerExecutor(dc dataChan) error {
	for {
		select {
		case i := <-dc:
			fmt.Printf("e is: %v \r\n", i)
		default:
			return nil
		}
	}
}
