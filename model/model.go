package model

import (
	"encoding/json"
	"math/rand"
	"time"
)

type GridPredInfo struct {
	StatisticDate    time.Time `gorm:"type:date;primary_key;column:STATISTIC_DATE" json:"statisticDate" `
	GridId           string    `gorm:"primary_key;column:GRID_ID" json:"gridId"`
	DairyGeneValue   float64   `gorm:"type:decimal(10,2);column:DAIRY_GENE_VALUE" json:"dairyGeneValue"`
	DairyLoadValue   float64   `sql:"type:decimal(10,2);column:DAIRY_LOAD_VALUE" json:"dairyLoadValue"`
	DairySalesValue  float64   `sql:"type:decimal(10,2);column:DAIRY_SALES_VALUE" json:"dairySalesValue"`
	DairyBuyValue    float64   `sql:"type:decimal(10,2);column:DAIRY_BUY_VALUE" json:"dairyBuyValue"`
	DairyIncomeValue float64   `sql:"type:decimal(10,4);column:DAIRY_INCOME_VALUE" json:"dairyIncomeValue"`
	DairySpendValue  float64   `sql:"type:decimal(10,4);column:DAIRY_SPEND_VALUE" json:"dairySpendValue"`
	PredGeneValue    float64   `sql:"type:decimal(10,2);column:PRED_GENE_VALUE" json:"predGeneValue"`
	PredBeamValue    float64   `sql:"type:decimal(10,2);column:PRED_BEAM_VALUE" json:"predBeamValue"`
	ActBeamValue     float64   `sql:"type:decimal(10,2);column:ACT_BEAM_VALUE" json:"actBeamValue"`
	PredRate         float64   `sql:"type:decimal(10,2);column:PRED_RATE" json:"predRate"`
	PredLoadValue    float64   `sql:"type:decimal(10,2);column:PRED_LOAD_VALUE" json:"predLoadValue"`
	AdjustValue      float64   `sql:"type:decimal(10,2);column:ADJUST_VALUE" json:"adjustValue"`
	AdjustDate       time.Time `sql:"column:ADJUST_DATE" json:"adjustDate"`
	AdjustUserId     string    `sql:"column:ADJUST_USER_ID" json:"adjustUserId"`
	CreateDate       time.Time `sql:"column:CREATE_DATE" json:"createDate"`
	UpdateDate       time.Time `sql:"column:UPDATE_DATE" json:"updateDate"`
}

func FindOne() {
	var v GridPredInfo
	MySqlDB.Find(&v,"STATISTIC_DATE =?",time.Now().Add(24*time.Hour*-1).Format("2006-01-02"))
	data,err:=json.Marshal(&v)
	if err!=nil{
		Logger.Info("解析出错 ")
	}
	Logger.Info("更新后的数据是 %s ",string(data))
}
func UpdateData(){
	var v GridPredInfo
	MySqlDB.Model(&v).Where("STATISTIC_DATE =?",time.Now().Add(24*time.Hour*-1).Format("2006-01-02")).
		Update("DAIRY_GENE_VALUE",getDataVal()).
		Update("DAIRY_LOAD_VALUE",getDataVal()).
		Update("PRED_GENE_VALUE",getDataVal())
	FindOne()
}

func getDataVal()(v float64){
	data:=[]float64{
		23,24,25,26,18,29,30,32,17,27,15,16,19,24,
	}
	//获取数组中的 随机数
	return data[RandInt64(1,12)]
}

func RandInt64(min, max int64) int64 {
	//以纳秒为随机种子
	rand.Seed(time.Now().UnixNano())
	//rand.Int63n(100) 生成100以内的随机数
	randNum := rand.Int63n(max - min) + min
	//休眠1纳秒，避免多次调用时,产生相同的int64值
	time.Sleep(time.Nanosecond)
	return randNum
}