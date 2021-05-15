package logic
type ResData struct {
	X []string
	Ori []string
	Real []string
	Predict []string
}
func Prediction(data []CodeList) ResData  {
	t:=ResData{}
	for _,v:=range data{
		t.X = append(t.X,v.Date)
	}
	for _,v:=range data{
		t.Real = append(t.Real,v.ClosePrice)
	}
	return t
}

func ekf(data []CodeList) {
	if len(data)<6{
		return
	}
	t:=ekf_t{}

	//初始化观测值
	t.x=data[0].convertSliceFloat()

	//量测矩阵为单位矩阵
	mat_addeye(t.H,4) //4维观测点


	//trainSeq:= len(data)*2/3
	//preditionSeq:= len(data)-trainSeq




	//ekf_init()
}