package logic

import (
	"fmt"
	"github.com/rosshemsley/kalman"
	"github.com/rosshemsley/kalman/models"
	"gonum.org/v1/gonum/mat"
	"time"
)

func Ekf(values [][]float64,iv,pv,ov float64)  (ori,predict map[int][]float64,newValue map[int]float64,){

	trainSeq:= len(values)*2/3
	preditionSeq:= len(values)-trainSeq
	ori=make(map[int][]float64,0)
	predict=make(map[int][]float64,0)
	newValue=make(map[int]float64,0)

	T:=time.Hour*5
	var t time.Time = time.Unix(0,0)
	data:=mat.NewVecDense(4,values[0])

	model:=models.NewBrownianModel(t,data,models.BrownianModelConfig{
		InitialVariance:     iv, //1.0,
		ProcessVariance:     pv, //1.0,
		ObservationVariance: ov, //2.0,
	})



	//model.TransitionRaw(1).Apply(func(i, j int, v float64) float64 {
	//	return 1
	//},model.TransitionRaw(1))


	//mat.DenseCopyOf(model.Transition(1)).Apply(func(i, j int, v float64) float64 {
	//	return 1.0
	//	if i==j{
	//		return 1.0
	//	}else{
	//		return 0.0
	//	}
	//},model.Transition(1))
	for j:=0;j<4;j++{
		ori[j]=append(ori[j],values[0][j])
	}

	filter := kalman.NewKalmanFilter(model)

	//_data:=mat.NewVecDense(4,values[0])
	//t_model:= model.NewMeasurement(_data)

	for i:=0;i< len(values);i++ {


		_data:=mat.NewVecDense(4,values[i])
		if i<=preditionSeq{//使用观测值修正结果
			for j:=0;j<4;j++{
				ori[j]=append(ori[j],values[i][j])
			}
		}else{//使用预测值继续预测
			var temp []float64
			for j:=0;j<4;j++{
				g:=filter.State().At(j,0)
				temp=append(temp,g)
			}
			//_data=mat.NewVecDense(4,temp)
			_data=mat.NewVecDense(4,values[i])
		}
		t_model:= model.NewMeasurement(_data)

		t:=filter.Time().Add(T)
		//预测
		err := filter.Predict(t)
		if err != nil {
			fmt.Println("err=",err.Error())
		}

		err = filter.Update(t,t_model)
		if err != nil {
			fmt.Println("err=",err.Error())
		}
		fmt.Printf("filtered value: %v\n", filter.State())

		if i==len(values)-1{
			//记录下一次的值
			for j:=0;j<4;j++{
				newValue[j]=filter.State().At(j,0)
			}
		}else{
			//记录kalman滤波结果
			for j:=0;j<4;j++{
				g:=filter.State().At(j,0)
				predict[j]=append(predict[j],g)
			}
		}


		//fmt.Printf("filtered value: %v\n", model.Value(filter.State()))

		//t_model:= model.NewPositionMeasurement(_data,0)
	}

	//t=filter.Time().Add(T)
	//err := filter.Predict(t)
	//if err != nil {
	//	fmt.Println("err=",err.Error())
	//}

	return
}