package show

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/iooikaak/frame/xlog"
)

const (
	LockPrefixShow     = "biz_show_lock_%d"
	TimeoutRepeatTimes = 20
)

func (d *Dao) GetShowLock(ctx context.Context, id int64) (locked bool, err error) {
	start := time.Now()
	repeat := TimeoutRepeatTimes
	for {
		locked, err = d.LockShow(ctx, id)
		if err != nil {
			s := fmt.Sprintf("%v", err)
			if strings.Contains(s, "timeout") && repeat > 0 {
				time.Sleep(100 * time.Millisecond)
				repeat -= 1
				continue
			}
			log.Errorf("failed to lock show id: %v err: %v", id, err)
			return locked, err
		}
		if locked {
			break
		}
		if time.Since(start) >= 2*time.Minute {
			return locked, fmt.Errorf("failed to lock show after 2 minutes id: %d", id)
		}
		time.Sleep(40 * time.Millisecond)
	}
	return locked, nil
}

func (d *Dao) LockShow(ctx context.Context, id int64) (res bool, err error) {
	var (
		key = fmt.Sprintf(LockPrefixShow, id)
	)
	res, err = d.redis.SetNX(ctx, key, "1", 300*time.Second).Result()
	if err != nil {
		return
	}
	if !res {
		return false, nil
	}
	return true, nil
}

func (d *Dao) ReleaseShowLock(ctx context.Context, id int64) (err error) {
	start := time.Now()
	repeat := TimeoutRepeatTimes
	for {
		err = d.UnLockShow(ctx, id)
		if err != nil {
			s := fmt.Sprintf("%v", err)
			if strings.Contains(s, "timeout") && repeat > 0 {
				time.Sleep(100 * time.Millisecond)
				repeat -= 1
				continue
			}
			if time.Since(start) >= 2*time.Minute {
				return fmt.Errorf("failed to release show after 2 minutes id: %d", id)
			}
			time.Sleep(40 * time.Millisecond)
			return err
		}
		break

	}
	return nil
}

func (d *Dao) UnLockShow(ctx context.Context, id int64) (err error) {
	var (
		key = fmt.Sprintf(LockPrefixShow, id)
	)
	_, err = d.redis.Del(ctx, key).Result()
	if err != nil {
		return
	}
	return
}
