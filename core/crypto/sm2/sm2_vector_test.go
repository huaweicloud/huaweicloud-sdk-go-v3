package sm2

import (
	"encoding/hex"
	"math/big"
	"testing"
)

type deterministicReader struct {
	pattern []byte
	pos     int
}

func (r *deterministicReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = r.pattern[r.pos%len(r.pattern)]
		r.pos++
	}
	return len(p), nil
}

var (
	sm2TestPrivD, _ = new(big.Int).SetString(
		"128B2FA8BD433C6C068C8D803DFF79792CF9F6FC0DAC5F89F8D97B8F3D5A8D6A", 16)
	sm2TestPubX, _ = new(big.Int).SetString(
		"1cb5ce0176dd2c6e441f1fc0868f6ec41fcc673c8f48f3dcad0a4d6fdfbacaeb", 16)
	sm2TestPubY, _ = new(big.Int).SetString(
		"f1c06cae78148eda3efed3b8a564e06e8cfbe35969097efe02e4451da27d88d5", 16)
	sm2TestR, _ = new(big.Int).SetString(
		"b8d073149ebc4a4dfd84d005f5fadd53e1d5749c32a3a850a9e86aa6ee7607f6", 16)
	sm2TestS, _ = new(big.Int).SetString(
		"4c533921cee71506956c0f6eefb1c79397f3373847e9d133d013a1535ed02573", 16)
	sm2TestMsg = []byte("hello sm2 deterministic test")
)

func newTestPrivateKey() *PrivateKey {
	curve := P256Sm2()
	priv := new(PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = sm2TestPrivD
	priv.PublicKey.X = sm2TestPubX
	priv.PublicKey.Y = sm2TestPubY
	return priv
}

func newDeterministicReader() *deterministicReader {
	pattern := make([]byte, 32)
	for i := range pattern {
		pattern[i] = byte(i)
	}
	return &deterministicReader{pattern: pattern}
}

func TestSm2DeterministicSign(t *testing.T) {
	priv := newTestPrivateKey()
	reader := newDeterministicReader()

	r, s, err := Sm2Sign(priv, sm2TestMsg, nil, reader)
	if err != nil {
		t.Fatalf("Sm2Sign failed: %v", err)
	}

	if r.Cmp(sm2TestR) != 0 {
		t.Errorf("r = %s, want %s", r.Text(16), sm2TestR.Text(16))
	}
	if s.Cmp(sm2TestS) != 0 {
		t.Errorf("s = %s, want %s", s.Text(16), sm2TestS.Text(16))
	}
}

func TestSm2VerifyDeterministicSignature(t *testing.T) {
	priv := newTestPrivateKey()

	ok := Sm2Verify(&priv.PublicKey, sm2TestMsg, nil, sm2TestR, sm2TestS)
	if !ok {
		t.Fatal("Sm2Verify should accept the deterministic signature")
	}

	wrongMsg := []byte("tampered message")
	ok = Sm2Verify(&priv.PublicKey, wrongMsg, nil, sm2TestR, sm2TestS)
	if ok {
		t.Fatal("Sm2Verify should reject signature for wrong message")
	}

	tamperedR := new(big.Int).Add(sm2TestR, big.NewInt(1))
	ok = Sm2Verify(&priv.PublicKey, sm2TestMsg, nil, tamperedR, sm2TestS)
	if ok {
		t.Fatal("Sm2Verify should reject tampered r")
	}

	tamperedS := new(big.Int).Add(sm2TestS, big.NewInt(1))
	ok = Sm2Verify(&priv.PublicKey, sm2TestMsg, nil, sm2TestR, tamperedS)
	if ok {
		t.Fatal("Sm2Verify should reject tampered s")
	}
}

func TestSm2PublicKeyVerifyDeterministic(t *testing.T) {
	priv := newTestPrivateKey()

	sig, err := SignDigitToSignData(sm2TestR, sm2TestS)
	if err != nil {
		t.Fatalf("SignDigitToSignData failed: %v", err)
	}

	ok := priv.PublicKey.Verify(sm2TestMsg, sig)
	if !ok {
		t.Fatal("PublicKey.Verify should accept the deterministic signature")
	}
}

func TestSm2DeterministicRoundTrip(t *testing.T) {
	priv := newTestPrivateKey()
	reader := newDeterministicReader()

	r, s, err := Sm2Sign(priv, sm2TestMsg, nil, reader)
	if err != nil {
		t.Fatalf("Sm2Sign failed: %v", err)
	}

	sig, err := SignDigitToSignData(r, s)
	if err != nil {
		t.Fatalf("SignDigitToSignData failed: %v", err)
	}

	ok := priv.PublicKey.Verify(sm2TestMsg, sig)
	if !ok {
		t.Fatal("Round-trip sign -> Verify failed")
	}

	r2, s2, err := SignDataToSignDigit(sig)
	if err != nil {
		t.Fatalf("SignDataToSignDigit failed: %v", err)
	}

	if r.Cmp(r2) != 0 || s.Cmp(s2) != 0 {
		t.Fatal("Sign digit round-trip conversion mismatch")
	}

	if r.Cmp(sm2TestR) != 0 || s.Cmp(sm2TestS) != 0 {
		t.Errorf("signature mismatch: r=%s s=%s", r.Text(16), s.Text(16))
	}
}

func TestSm2KeyDerivationConsistency(t *testing.T) {
	curve := P256Sm2()

	priv := new(PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = sm2TestPrivD
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(sm2TestPrivD.Bytes())

	if priv.PublicKey.X.Cmp(sm2TestPubX) != 0 {
		t.Errorf("public key X = %s, want %s",
			hex.EncodeToString(priv.PublicKey.X.Bytes()),
			hex.EncodeToString(sm2TestPubX.Bytes()))
	}
	if priv.PublicKey.Y.Cmp(sm2TestPubY) != 0 {
		t.Errorf("public key Y = %s, want %s",
			hex.EncodeToString(priv.PublicKey.Y.Bytes()),
			hex.EncodeToString(sm2TestPubY.Bytes()))
	}
}
