package interfaces

type HV struct {
	canh float64
}

func (hv *HV) chuVi() float64 {
	return hv.canh * 4
}

func (hv *HV) dienTich() float64 {
	return hv.canh * hv.canh
}
