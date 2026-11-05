package sm2

import (
	"crypto/rand"
	"sync"
	"testing"
)

func TestSm2ConcurrentSignVerify(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	msg := []byte("concurrent test message")
	var wg sync.WaitGroup
	concurrency := 100

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, s, err := Sm2Sign(priv, msg, nil, rand.Reader)
			if err != nil {
				t.Errorf("Sm2Sign failed: %v", err)
				return
			}
			if !Sm2Verify(&priv.PublicKey, msg, nil, r, s) {
				t.Error("Sm2Verify failed in concurrent context")
			}
		}()
	}
	wg.Wait()
}

func TestSm2ConcurrentP256Init(t *testing.T) {
	var wg sync.WaitGroup
	concurrency := 50

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			curve := P256Sm2()
			if curve == nil {
				t.Error("P256Sm2 returned nil")
			}
			if curve.Params().BitSize != 256 {
				t.Error("wrong BitSize")
			}
		}()
	}
	wg.Wait()
}
