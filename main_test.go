package randominfo

import "testing"

func TestRandomPhone(t *testing.T) {
	phone, err := RandomPhone()
	if err != nil {
		t.Errorf("err = %v; expected nil", err)
	}
	if len(phone) != 11 {
		t.Errorf("phone = %v; expected 11 length", phone)
	}

}

func TestRandomIDcard(t *testing.T) {
	id, err := RandomIDcard("150701")
	if err != nil {
		t.Errorf("err = %v; expected nil", err)
	}
	if len(id) != 18 {
		t.Errorf("id = %v; expected 11 length", id)
	}
}
func TestRandomInfo(t *testing.T) {
	info, err := RandomInfo()
	if err != nil {
		t.Errorf("err = %v; expected nil", err)
	}
	t.Log(info)
}
