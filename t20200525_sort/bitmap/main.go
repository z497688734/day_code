package main

import "fmt"

func main()  {




	arrInt32 := [...]uint32{5, 4, 2, 1, 3, 17, 13}

	var arrMax uint32 = 17
	bit := NewBitmap(arrMax)

	for _, v := range arrInt32 {
		bit.Set(v)
	}

	fmt.Println("排序后:")
	for i := uint32(0); i < arrMax; i++ {
		//fmt.Println("test i",i,"ret",bit.Test(i))


		if k := bit.Test(i); k == true{
			fmt.Printf("%d ", i)
		}
	}



	//bit := NewBitmap(17)
	//bit.Set(2)
	//fmt.Println(bit.Test(2))
}


const (
	BitSize = 8 //一个字节8位
)


type Bitmap struct {
	BitArray  []byte
	ArraySize uint32
}


//一个byte有8位,可代表8个数字,取余后加1为存放最大数所需的容量
func NewBitmap(max uint32) *Bitmap {
	var r uint32
	switch {
	case max <= BitSize:
		r = 1
	default:
		r = max / BitSize
		if max%BitSize != 0 {
			r += 1
		}

		fmt.Println(r)
	}

	fmt.Println("数组大小:", r)
	return &Bitmap{BitArray: make([]byte, r), ArraySize: r}
}



func (bitmap *Bitmap) Set(i uint32) {
	idx, pos := bitmap.calc(i)
	fmt.Println(1 << pos)
	bitmap.BitArray[idx] |= 1 << pos  // |= 表示已经存在的不能改变
	fmt.Println("set()  value=", i, " idx=", idx, " pos=", pos, ByteToBinaryString(bitmap.BitArray[idx]))
}

func (bitmap *Bitmap) Test(i uint32) bool {
	idx, pos := bitmap.calc(i)
	return (bitmap.BitArray[idx] & (1<< pos) !=0)
}

func (bitmap *Bitmap) Clear(i uint32) {
	idx, pos := bitmap.calc(i)
	bitmap.BitArray[idx] &^= 1 << pos
}

func (bitmap *Bitmap) calc(i uint32) (idx, pos uint32) {

	idx = i >> 3 //相当于i / 8,即字节位置
	if idx >= bitmap.ArraySize {
		panic("数组越界.")
		return
	}
	pos = i % BitSize //位位置
	return
}


func ByteToBinaryString(data byte) (str string) {
	var a byte
	for i := 0; i < 8; i++ {
		a = data
		data <<= 1
		data >>= 1

		switch a {
		case data:
			str += "0"
		default:
			str += "1"
		}



		data <<= 1
	}
	return str
}