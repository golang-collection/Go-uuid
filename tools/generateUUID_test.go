package tools

import "testing"

/**
* @Author: super
* @Date: 2020-09-09 20:35
* @Description:
**/
func TestGenerateUUIDV1(t *testing.T) {
	uuid := GenerateUUIDV1()
	t.Log(uuid)
}

func TestGenerateUUIDV2(t *testing.T) {
	uuid, _ := GenerateUUIDV2()
	t.Log(uuid)
}

func TestGenerateUUIDV3(t *testing.T) {
	uuid := GenerateUUIDV3()
	t.Log(uuid)
}

func TestGenerateUUIDV4(t *testing.T) {
	uuid := GenerateUUIDV4()
	t.Log(uuid)
}

func BenchmarkGenerateUUIDV4(b *testing.B) {
	for i := 0; i<b.N; i++{
		_ = GenerateUUIDV4()
	}
}

func TestGenerateUUIDV5(t *testing.T) {
	uuid := GenerateUUIDV5()
	t.Log(uuid)
}