package interfaces

type HCN struct {
	dai  float64
	rong float64
}

func (hcn *HCN) chuVi() float64 {
	return (hcn.dai + hcn.rong) / 2
}

func (hcn *HCN) dienTich() float64 {
	return hcn.dai * hcn.rong
}
