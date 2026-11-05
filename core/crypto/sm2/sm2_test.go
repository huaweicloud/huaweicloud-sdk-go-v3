package sm2

import (
	"crypto/rand"
	"encoding/hex"
	"testing"
)

func TestP256Sm2Curve(t *testing.T) {
	curve := P256Sm2()
	if curve == nil {
		t.Fatal("P256Sm2() returned nil")
	}
	params := curve.Params()
	if params.BitSize != 256 {
		t.Errorf("BitSize = %d, want 256", params.BitSize)
	}
	if params.Name != "SM2-P-256" {
		t.Errorf("Name = %s, want SM2-P-256", params.Name)
	}
}

func TestGenerateKey(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}
	if priv == nil {
		t.Fatal("GenerateKey returned nil")
	}
	if priv.D == nil || priv.D.Sign() <= 0 {
		t.Fatal("Private key D is invalid")
	}
	if priv.X == nil || priv.Y == nil {
		t.Fatal("Public key X or Y is nil")
	}
	if !priv.Curve.IsOnCurve(priv.X, priv.Y) {
		t.Fatal("Public key is not on curve")
	}
}

func TestSm2SignVerify(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	msg := []byte("hello sm2 sign verify test")
	r, s, err := Sm2Sign(priv, msg, nil, rand.Reader)
	if err != nil {
		t.Fatalf("Sm2Sign failed: %v", err)
	}
	if r.Sign() <= 0 || s.Sign() <= 0 {
		t.Fatal("Sign returned non-positive values")
	}

	ok := Sm2Verify(&priv.PublicKey, msg, nil, r, s)
	if !ok {
		t.Fatal("Sm2Verify failed for valid signature")
	}

	wrongMsg := []byte("wrong message")
	ok = Sm2Verify(&priv.PublicKey, wrongMsg, nil, r, s)
	if ok {
		t.Fatal("Sm2Verify should fail for wrong message")
	}
}

func TestSm2EncryptDecrypt(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	plaintext := []byte("hello sm2 encrypt decrypt test")
	ciphertext, err := Encrypt(&priv.PublicKey, plaintext, rand.Reader, C1C3C2)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(priv, ciphertext, C1C3C2)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if hex.EncodeToString(decrypted) != hex.EncodeToString(plaintext) {
		t.Fatalf("Decrypted = %x, want %x", decrypted, plaintext)
	}
}

func TestSm2EncryptDecryptC1C2C3(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	plaintext := []byte("hello sm2 C1C2C3 test")
	ciphertext, err := Encrypt(&priv.PublicKey, plaintext, rand.Reader, C1C2C3)
	if err != nil {
		t.Fatalf("Encrypt C1C2C3 failed: %v", err)
	}

	decrypted, err := Decrypt(priv, ciphertext, C1C2C3)
	if err != nil {
		t.Fatalf("Decrypt C1C2C3 failed: %v", err)
	}

	if hex.EncodeToString(decrypted) != hex.EncodeToString(plaintext) {
		t.Fatalf("Decrypted = %x, want %x", decrypted, plaintext)
	}
}

func TestSm2EncryptDecryptAsn1(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	plaintext := []byte("hello sm2 asn1 test")
	ciphertext, err := EncryptAsn1(&priv.PublicKey, plaintext, rand.Reader)
	if err != nil {
		t.Fatalf("EncryptAsn1 failed: %v", err)
	}

	decrypted, err := DecryptAsn1(priv, ciphertext)
	if err != nil {
		t.Fatalf("DecryptAsn1 failed: %v", err)
	}

	if hex.EncodeToString(decrypted) != hex.EncodeToString(plaintext) {
		t.Fatalf("Decrypted = %x, want %x", decrypted, plaintext)
	}
}

func TestPublicKeyVerify(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	msg := []byte("test public key verify")
	sig, err := priv.Sign(rand.Reader, msg, nil)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	ok := priv.PublicKey.Verify(msg, sig)
	if !ok {
		t.Fatal("PublicKey.Verify failed for valid signature")
	}
}

func TestSm3Digest(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	msg := []byte("test sm3 digest")
	digest, err := priv.PublicKey.Sm3Digest(msg, nil)
	if err != nil {
		t.Fatalf("Sm3Digest failed: %v", err)
	}
	if len(digest) != 32 {
		t.Fatalf("Digest length = %d, want 32", len(digest))
	}
}

func TestIsOnCurve(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}
	if !P256Sm2().IsOnCurve(priv.X, priv.Y) {
		t.Fatal("Generated key point is not on curve")
	}
}

func TestSignDigitToSignData(t *testing.T) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}

	msg := []byte("test sign digit conversion")
	r, s, err := Sm2Sign(priv, msg, nil, rand.Reader)
	if err != nil {
		t.Fatalf("Sm2Sign failed: %v", err)
	}

	signData, err := SignDigitToSignData(r, s)
	if err != nil {
		t.Fatalf("SignDigitToSignData failed: %v", err)
	}

	r2, s2, err := SignDataToSignDigit(signData)
	if err != nil {
		t.Fatalf("SignDataToSignDigit failed: %v", err)
	}

	if r.Cmp(r2) != 0 || s.Cmp(s2) != 0 {
		t.Fatal("Sign digit round-trip conversion mismatch")
	}
}
