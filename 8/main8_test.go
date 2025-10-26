package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup_Basic(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer wg.Close()

	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(20 * time.Millisecond)
	}()

	start := time.Now()
	wg.Wait()
	duration := time.Since(start)

	if duration < 20*time.Millisecond {
		t.Errorf("Wait завершился слишком рано: %v", duration)
	}
}

func TestCustomWaitGroup_Zero(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer wg.Close()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// OK
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait не завершился немедленно для пустой WaitGroup")
	}
}

func TestCustomWaitGroup_NegativeCounter(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer wg.Close()

	// Вызываем Done() без предварительного Add()
	// Счётчик должен остаться на 0
	wg.Done()

	// Проверяем, что Wait() завершается немедленно
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// OK - Wait завершился немедленно
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait не завершился немедленно после отрицательного счётчика")
	}
}

func TestCustomWaitGroup_MultipleWaits(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer wg.Close()

	wg.Add(1)

	var done1, done2 chan struct{}
	go func() {
		wg.Wait()
		done1 = make(chan struct{})
		close(done1)
	}()

	go func() {
		wg.Wait()
		done2 = make(chan struct{})
		close(done2)
	}()

	time.Sleep(10 * time.Millisecond)

	wg.Done()

	select {
	case <-done1:
	case <-time.After(100 * time.Millisecond):
		t.Error("Первый Wait не разблокировался")
	}

	select {
	case <-done2:
	case <-time.After(100 * time.Millisecond):
		t.Error("Второй Wait не разблокировался")
	}
}

func TestCustomWaitGroup_AddAfterWait(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer wg.Close()

	wg.Add(1)

	doneWait := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneWait)
	}()

	time.Sleep(10 * time.Millisecond)

	wg.Add(1)

	wg.Done()

	select {
	case <-doneWait:
		t.Error("Wait завершился раньше времени")
	default:
	}

	wg.Done()

	select {
	case <-doneWait:
		// OK
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait не завершился после второй Done")
	}
}
