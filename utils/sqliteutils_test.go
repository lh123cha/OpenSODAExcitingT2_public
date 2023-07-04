package utils

import (
	"strconv"
	"testing"
)

func TestInsertsinglequery(t *testing.T) {
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3.0, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}

	testQueryresult := make(map[string](map[string]interface{}))

	testQueryresult["openrank"] = testData1
	//test_repoinfo := service.GetRepoInfoOfMetric("X-lab2017/open-digger", "openrank")

	err := Insertsinglequery("open-digger", "https://oss.x-lab.info/open_digger/github/X-lab2017/open-digger", testQueryresult)
	if err != nil {
		t.Fatal("Insert error!")
	}
}

func TestReadquery(t *testing.T) {
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3.0, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testQueryresult := make(map[string](map[string]interface{}))
	metric := "openrank"
	testQueryresult[metric] = testData1

	res := Repoinfo{}
	Readquery(&res, "open-digger")
	for k, v := range testData1 {
		if res.Data[metric][k].(float64) != v {
			t.Errorf("Read query error" + k + strconv.FormatFloat(float64(res.Data[metric][k].(float64)), 'f', 6, 64) + " correct:" + strconv.FormatFloat(v.(float64), 'f', 6, 64))
		}
	}
}
func TestInsertlog(t *testing.T) {
	testlog1 := "opendigger repo = X-lab2017/open-digger metric = OpenRank month = 2023-01"
	testlog2 := "opendigger repo = X-lab2017/open-digger metric = OpenRank"
	testlog3 := "opendigger repo = X-lab2017/open-digger month = 2023-01"
	err := Insertlog(testlog1)
	err = Insertlog(testlog2)
	err = Insertlog(testlog3)
	if err != nil {
		t.Fatal("Insert log error")
	}
}
func TestReadlog(t *testing.T) {
	logs := []Searchhistory{}
	Readlog(&logs)
	testlogs := []string{
		"opendigger repo = X-lab2017/open-digger metric = OpenRank month = 2023-01",
		"opendigger repo = X-lab2017/open-digger metric = OpenRank",
		"opendigger repo = X-lab2017/open-digger month = 2023-01",
	}
	for i, log := range testlogs {
		if logs[i].Log != log {
			t.Errorf("Read logs error " + "\"" + logs[i].Log + "\"" + " Correct is :" + "\"" + log + "\"")
		}
	}
}
