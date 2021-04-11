package tripleDES

import "testing"

func TestTripleDES(t *testing.T) {

	tests := []struct {
		name       string
		a          string
		K1, K2, K3 string
	}{
		{
			name: "test-1",
			a:    "Go Lang is Awesome",
			K1:   "12312312",
			K2:   "34343434",
			K3:   "56565656",
		},
		{
			name: "test-2",
			a:    "ƕ��ֆ�W%�3��3�C�g��'��j8�",
			K1:   "12312312",
			K2:   "34343434",
			K3:   "56565656",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Before encryption:%_v", tt.a)
			triple_keys := tt.K1 + tt.K2 + tt.K3
			encrypted, err := TripleDESEncryption([]byte(tt.a), []byte(triple_keys))
			if err != nil {
				panic(err)
			}
			t.Log("After encryption:%_v", string(encrypted[:]))

			decrypted, err := TripleDESDecryption(encrypted, []byte(triple_keys))
			if err != nil {
				panic(err)
			}

			t.Log("After decryption:%_v", string(decrypted[:]))

			if tt.a == string(decrypted) {
				t.Log("Plain Text is encrypted and decrypted correct")
			} else {
				t.Log("Wrong Encryption")
			}
		})
	}

}
