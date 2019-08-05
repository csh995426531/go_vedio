package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/csh995426531/go_vedio/scheduler/dbops"
)

// VideoClearDispatcher 生产待处理任务插入data chan
func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}
	return nil
}

// VideoClearExecutor 消费待处理任务data chan
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	forloop:
		for {
			select {
			case vid := <-dc:
				go func(id string) {
					if err := deleteVideo(id); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id); err != nil {
						errMap.Store(id, err)
						return
					}
					log.Printf("del video id: %v successful", id)
				}(vid.(string))
			default:
				break forloop
			}
		}

	errMap.Range(func(k, v interface{}) bool {
		err := v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}

// deleteVideo 删除视频文件
func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}
