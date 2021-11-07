package zhihuishu

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"
	"zhs/zhihuishu/errorCode"
	"zhs/zhihuishu/util"
)

// 查询所有一级课程
func queryShareCourseInfo(uuid string) (shareCourseInfo ShareCourseInfoStruct, err error) {
	api := "https://onlineservice.zhihuishu.com/student/course/share/queryShareCourseInfo"
	// 构造payload
	payload := url.Values{}
	payload.Set("status", "0")
	payload.Set("pageNo", "1")
	payload.Set("pageSize", "5")
	payload.Set("uuid", uuid)
	payload.Set("date", time.Now().UTC().String())
	header["Cookie"] = cookie1
	curlShareCourseInfo, curlErr := util.Curl("POST", api, payload.Encode(), header)
	if curlErr != nil {
		return ShareCourseInfoStruct{}, curlErr
	}
	if jsonParseErr := json.Unmarshal(curlShareCourseInfo, &shareCourseInfo); jsonParseErr != nil {
		return ShareCourseInfoStruct{}, fmt.Errorf("jsonParseErr: %s, 错误代码：%d ",jsonParseErr, errorCode.JsonUnMarshalError)
	}
	if shareCourseInfo.Code != 200 || len(shareCourseInfo.Result.CourseOpenDtos) == 0 {
		return ShareCourseInfoStruct{}, fmt.Errorf("登录已过期！")
	}
	return shareCourseInfo, nil
}

// 获取视频列表
func getVideoList(uuid, recruitAndCourseId string) (VideoListStruct , error) {
	header["Cookie"] = cookie2
	api := "https://studyservice.zhihuishu.com/learning/videolist"
	payload := url.Values{}
	payload.Set("recruitAndCourseId", recruitAndCourseId)
	payload.Set("uuid", uuid)
	payload.Set("dateFormate", util.GetStringTimestamp())
	queryVideoList, curlErr := util.Curl("POST", api, payload.Encode(), header)
	if curlErr != nil {
		return VideoListStruct{}, curlErr
	}
	videoList := VideoListStruct{}
	if parseJsonErr := json.Unmarshal(queryVideoList, &videoList); parseJsonErr != nil {
		return VideoListStruct{}, parseJsonErr
	}
	if videoList.Code != 0 {
		return VideoListStruct{}, fmt.Errorf("错误信息：%s ", videoList.Message)
	}
	return videoList, nil
}

// 获取之前的学习信息
func preLearningNote(courseId, chapterId, lessonId, recruitId, videoId int, uuid string) (PreLearningNoteStruct, error)  {
	header["Cookie"] = cookie2
	api := "https://studyservice.zhihuishu.com/learning/prelearningNote"
	payload := url.Values{}
	payload.Set("ccCourseId", strconv.FormatInt(int64(courseId), 10))
	payload.Set("chapterId", strconv.FormatInt(int64(chapterId), 10))
	payload.Set("isApply", "1")
	payload.Set("lessonId", strconv.FormatInt(int64(lessonId), 10))
	payload.Set("recruitId", strconv.FormatInt(int64(recruitId), 10))
	payload.Set("videoId", strconv.FormatInt(int64(videoId), 10))
	payload.Set("uuid", uuid)
	payload.Set("dateFormate", util.GetStringTimestamp())
	preLearningNoteByte, curlErr := util.Curl("POST", api, payload.Encode(), header)
	if curlErr != nil {
		return PreLearningNoteStruct{}, curlErr
	}
	preLearningNoteData := PreLearningNoteStruct{}
	if jsonParseErr := json.Unmarshal(preLearningNoteByte, &preLearningNoteData); jsonParseErr != nil {
		return PreLearningNoteStruct{}, jsonParseErr
	}
	if preLearningNoteData.Code != 0 {
		return PreLearningNoteStruct{}, fmt.Errorf("错误信息：%s ", preLearningNoteData.Message)
	}
	return preLearningNoteData, nil
}

// 获取Ev
func getEv(param map[string]string) string {
	payload := GetEvReqJsonStruct{
		RecruitId: param["recruitId"],
		LessonId:  param["lessonId"],
		Small:     param["small"],
		VideoId:   param["videoId"],
		ChapterId: param["chapterId"],
		Status:    param["status"],
		PlaySec:   param["playSec"],
		TotalSec:  param["totalSec"],
		Position:  param["position"],
	}
	marshal, jsonStringifyErr := json.Marshal(payload)
	if jsonStringifyErr != nil {
		return ""
	}
	getEV, getEvErr := util.Curl("POST", "http://schoolnet.ufec.cn/zhihuishu", string(marshal), map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
	if getEvErr != nil {
		fmt.Println("getEvErr: ", getEvErr)
		return ""
	}
	evRes := GetEvResStruct{}
	if jsonParseErr := json.Unmarshal(getEV, &evRes); jsonParseErr != nil {
		fmt.Println("jsonParseErr： ", jsonParseErr)
		return ""
	}
	return evRes.EV
}

// 保存学习记录
func saveStudyTime(param map[string]string)  {
	fmt.Println("开始学习了吗？ https://www.ufec.cn")
	payload := url.Values{}
	payload.Set("watchPoint", "")
	payload.Set("ev", getEv(param))
	payload.Set("learningTokenId", param["learningTokenId"])
	payload.Set("courseId", param["courseId"])
	payload.Set("uuid", param["uuid"])
	payload.Set("dateFormate", util.GetStringTimestamp())
	header["Cookie"] = cookie2
	saveStudyTimeCurl, saveStudyTimeErr := util.Curl("POST", "https://studyservice.zhihuishu.com/learning/saveDatabaseIntervalTime", payload.Encode(), header)
	if saveStudyTimeErr != nil {
		fmt.Println("saveStudyTimeErr:" ,saveStudyTimeErr)
		return
	}
	saveStudyTime := SaveStudyTimeStruct{}
	if jsonParseErr := json.Unmarshal(saveStudyTimeCurl, &saveStudyTime); jsonParseErr != nil {
		fmt.Println("jsonParseErr: ", jsonParseErr)
		return
	}
	if saveStudyTime.Code != 0 {
		log.Print(saveStudyTime.Message)
		return
	}else{
		log.Print(param["name"], "课程学习完成，前往智慧树官网或App查看")
	}
}