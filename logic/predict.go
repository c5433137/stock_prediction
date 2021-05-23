package logic

type ResData struct {
	Stock string
	X []int
	Ori map[int][]float64
	Real map[int][]float64
	Predict map[int][]float64
	NewValue map[int]float64
}
func Prediction(data []CodeListNewType,iv,pv,ov float64) ResData  {
	t:=ResData{
		Ori:make(map[int][]float64,0),
		Real:make(map[int][]float64,0),
		Predict:make(map[int][]float64,0),
		NewValue:make(map[int]float64,0),
	}
	for _,v:=range data{
		t.X = append(t.X,v.Date)
	}
	for _,v:=range data{
		t.Real[0] = append(t.Real[0],v.OpenPrice)
		t.Real[1] = append(t.Real[1],v.ClosePrice)
		t.Real[2] = append(t.Real[2],v.MinPrice)
		t.Real[3] = append(t.Real[3],v.MaxPrice)
	}
	t.Stock = data[0].Code
	t.Ori,t.Predict,t.NewValue = ekf(data,iv,pv,ov)
	return t
}

func ekf(data []CodeListNewType,iv,pv,ov float64) (ori,predict map[int][]float64,newValue map[int]float64,){
	if len(data)<6{
		return
	}
	var values [][]float64
	for _,v:=range data{
		values=append(values,v.convertSliceFloat())
	}
	return Ekf(values,iv,pv,ov)

}