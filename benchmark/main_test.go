package main

import (
	"testing"
)

func BenchmarkCubeTurnFaceBy(b *testing.B) {

	b.ResetTimer()
	b.Run("Cube", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}
			c = c.turnFace(1, 1)
		}
	})
	b.Run("Helper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}
			c = cube2x2x2Helper.turnFace(c, 1, 1)
		}
	})
	b.Run("HelperMod", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}
			c = cube2x2x2HelperMod.turnFace(c, 1, 1)
		}
	})
	b.Run("HelperModByte", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}
			c = cube2x2x2HelperModByte.turnFace(c, 1, 1)
		}
	})
}

func BenchmarkCubeGetID(b *testing.B) {

	b.ResetTimer()
	b.Run("Cube", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{
				49, 49, 49, 49,
				50, 50, 50, 50,
				52, 52, 52, 52,
				53, 53, 53, 53,
				48, 48, 48, 48,
				51, 51, 51, 51,
			}
			_ = c.getIDString()
		}
	})
	b.Run("Helper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{
				49, 49, 49, 49,
				50, 50, 50, 50,
				52, 52, 52, 52,
				53, 53, 53, 53,
				48, 48, 48, 48,
				51, 51, 51, 51,
			}
			_ = cube2x2x2Helper.getIDString(c)
		}
	})
	b.Run("CubeCustom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{
				49, 49, 49, 49,
				50, 50, 50, 50,
				52, 52, 52, 52,
				53, 53, 53, 53,
				48, 48, 48, 48,
				51, 51, 51, 51,
			}
			_ = c.getIDCustom()
		}
	})
	b.Run("HelperCustom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{
				49, 49, 49, 49,
				50, 50, 50, 50,
				52, 52, 52, 52,
				53, 53, 53, 53,
				48, 48, 48, 48,
				51, 51, 51, 51,
			}
			_ = cube2x2x2Helper.getIDCustom(&c)
		}
	})
	b.Run("JustString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := cube{
				49, 49, 49, 49,
				50, 50, 50, 50,
				52, 52, 52, 52,
				53, 53, 53, 53,
				48, 48, 48, 48,
				51, 51, 51, 51,
			}
			_ = string(c[:])
		}
	})
}

func BenchmarkSolve(b *testing.B) {
	c := cube{
		49, 49, 49, 49,
		50, 50, 50, 50,
		52, 52, 52, 52,
		53, 53, 53, 53,
		48, 48, 48, 48,
		51, 51, 51, 51,
	}
	id := string(c[:])

	e := engine2x2x2{
		permutations: map[string]*permutation{},
		ids:          []string{},
		helper:       &cube2x2x2HelperMod,
	}
	e.permutations[id] = &permutation{
		cube:  c,
		level: 0,
	}
	e.ids = append(e.ids, id)

	f := engine2x2x2{
		permutations: map[string]*permutation{},
		ids:          []string{},
		helper:       &cube2x2x2HelperMod,
	}
	f.permutations[id] = &permutation{
		cube:  c,
		level: 0,
	}
	f.ids = append(e.ids, id)

	loops := 0
	for len(e.ids) != loops {
		e.run(e.ids[loops], 16)
		loops++
	}
	b.ResetTimer()
	b.Run("engine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.run(e.ids[i], 16)
		}
	})
	b.Run("function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			run(&f, f.ids[i], 16)
		}
	})

}

func BenchmarkMapSearch(b *testing.B) {
	var i24Byte [24]byte
	mFloat32 := map[float32]string{}
	mString := map[string]string{}
	m24Byte := map[[24]byte]string{}

	idsFloat32 := []float32{}
	idsString := []string{}
	ids24Byte := [][24]byte{}

	for i := 0; i < 1000000; i++ {
		iString := string(i)
		iByte := []byte(iString)
		copy(i24Byte[:], iByte[:24])

		mFloat32[float32(i)] = iString
		mString[iString] = iString
		m24Byte[i24Byte] = iString

		idsFloat32 = append(idsFloat32, float32(i))
		idsString = append(idsString, iString)
		ids24Byte = append(ids24Byte, i24Byte)
	}

	b.ResetTimer()
	b.Run("Float32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range idsFloat32 {
				_, _ = mFloat32[v]
			}
		}
	})
	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range idsString {
				_, _ = mString[v]
			}
		}
	})
	b.Run("24Byte", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range ids24Byte {
				_, _ = m24Byte[v]
			}
		}
	})
}

func BenchmarkSliceInsert(b *testing.B) {
	var i24Byte [24]byte

	idsFloat32 := []float32{}
	idsString := []string{}
	ids24Byte := [][24]byte{}

	insertFloat32 := []float32{}
	insertString := []string{}
	insert24Byte := [][24]byte{}

	for i := 0; i < 1000000; i++ {
		iString := string(i)
		iByte := []byte(iString)
		copy(i24Byte[:], iByte[:24])

		idsFloat32 = append(idsFloat32, float32(i))
		idsString = append(idsString, iString)
		ids24Byte = append(ids24Byte, i24Byte)
	}

	b.ResetTimer()
	b.Run("Float32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range idsFloat32 {
				insertFloat32 = append(insertFloat32, v)
			}
		}
	})
	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range idsString {
				insertString = append(insertString, v)
			}
		}
	})
	b.Run("24Byte", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range ids24Byte {
				insert24Byte = append(insert24Byte, v)
			}
		}
	})
}
