package usecase

import (
	"fmt"

	"github.com/hanifmaliki/golang-lock/redis-lock/infrastructure"
)

type ProductUseCase struct {
	repo *infrastructure.ProductRepository
	lock *infrastructure.RedisLock
}

func (u *ProductUseCase) UpdateProductPrice(id int, newPrice float64) error {
	lockKey := fmt.Sprintf("product_lock:%d", id)

	// Try to acquire the lock
	locked, err := u.lock.AcquireLock(lockKey)
	if err != nil {
		return fmt.Errorf("error acquiring lock: %v", err)
	}
	if !locked {
		return fmt.Errorf("could not acquire lock, try again later")
	}

	// Ensure that the lock is released when the function exits
	defer func() {
		err := u.lock.ReleaseLock(lockKey)
		if err != nil {
			fmt.Printf("Error releasing lock: %v\n", err)
		}
	}()

	// Perform product price update
	err = u.repo.UpdateProductPrice(id, newPrice)
	if err != nil {
		return fmt.Errorf("error updating product price: %v", err)
	}

	return nil
}
