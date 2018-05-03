package redlock

import (
	"testing"
)

func TestRedLockSimple(t *testing.T) {

	c1, err := NewRedisHandler("localhost:6379")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	client1 := RedLock{
		C : c1,
		Password : 0,
		Key : "UniqueKey",
	}

	c2, err := NewRedisHandler("localhost:6379")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	client2 := RedLock{
		C : c2,
		Password : 0,
		Key : "UniqueKey",
	}

	ok := client1.AcquireLock(false) 
	if ok == false {
		t.Errorf("should Get the lock")
	}

	ok = client2.AcquireLock(false) 
	if ok == true {
		t.Errorf("should NOT Get the lock")
	}

}

func TestRedLockReleaseSimple(t *testing.T) {

	c1, err := NewRedisHandler("localhost:6379")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	client1 := RedLock{
		C : c1,
		Password : 0,
		Key : "UniqueKey1",
	}

	ok := client1.AcquireLock(false) 
	if ok == false {
		t.Errorf("should Get the lock")
	}

	ok = client1.ReleaseLock() 
	if ok == false {
		t.Errorf("should Release the lock")
	}

}