package syncgroup

import (
	"testing"
)

func testSyncGroup(t *testing.T, sg1, sg2 *SyncGroup) {
	n := 16
	exited := make(chan bool, n)

	for i := 0; i < n; i++ {
		sg1.Add()
		go func() {
			sg1.Done()
			sg2.Add()
			sg2.Wait()
			exited <- true
		}()
	}

	sg1.Wait()

	for i := 0; i < n; i++ {
		select {
		case <-exited:
			t.Fatal("WaitGroup released group too soon")
		default:
		}
		sg2.Done()
	}

	for i := 0; i < n; i++ {
		<-exited
	}
}

func TestSyncGroup(t *testing.T) {
	sg1 := New(8)
	sg2 := New(8)

	// Run the same test a few times to ensure barrier is in a proper state.
	for i := 0; i != 8; i++ {
		testSyncGroup(t, sg1, sg2)
	}
}
