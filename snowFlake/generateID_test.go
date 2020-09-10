package snowFlake

import "testing"

/**
* @Author: super
* @Date: 2020-09-09 22:07
* @Description:
**/

func TestGenerateSnowflake(t *testing.T) {
	err := InitSnowflake()
	if err != nil{
		t.Error(err)
		return
	}
	snowflake := GenerateSnowflake()
	t.Log(snowflake)
}

func BenchmarkGenerateSnowflake(b *testing.B) {
	err := InitSnowflake()
	if err != nil{
		b.Error(err)
		return
	}
	for i:=0; i<b.N; i++{
		_ = GenerateSnowflake()
	}
}