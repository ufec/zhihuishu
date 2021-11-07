package zhihuishu

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"time"
	"zhs/zhihuishu/util"
)

func Run() {
	fmt.Println("注意：本程序开源免费，如滥用或用于商业，随时关闭")
	fmt.Println("请输入uuid：")
	fmt.Scanln(&uuid)
	fmt.Println("请输入 onlineweb 页面 cookie：")
	fmt.Scanln(&cookie1)
	fmt.Println("请输入 studyh5 页面Cookie：")
	fmt.Scanln(&cookie2)
	fmt.Println("是否需要等待（等待将以1.5倍速模拟观看，可能需要较长时间）等待输入1，不等待输入0：")
	fmt.Scanln(&isNeed)
	fmt.Println("uuid为：", uuid)
	fmt.Println("cookie1为", cookie1)
	fmt.Println("cookie2为", cookie2)
	running()
}

func running()  {
	// 获取公开课信息
	shareCourseInfo, queryShareCourseInfoErr := queryShareCourseInfo(uuid)
	if queryShareCourseInfoErr != nil {
		fmt.Println(queryShareCourseInfoErr)
		return
	}
	// 遍历所有公开课
	for _, shareCourseDto := range shareCourseInfo.Result.CourseOpenDtos {
		// 获取课程下视频列表
		videoList, getVideoListErr := getVideoList(uuid, shareCourseDto.Secret)
		if getVideoListErr != nil {
			fmt.Println("getVideoListErr：", getVideoListErr)
			return
		}
		for _, chapter := range videoList.Data.VideoChapterDtos {
			for _, lesson := range chapter.VideoLessons {
				tempMap := make(map[string]string, 10)
				// 查看当前课程之前的学习记录
				preLearningNote, getPreLearningNoteErr := preLearningNote(videoList.Data.CourseID, lesson.ChapterID, lesson.ID, videoList.Data.RecruitID, lesson.VideoID, uuid)
				if getPreLearningNoteErr != nil {
					fmt.Println("\t\t", getPreLearningNoteErr)
					return
				}
				// 没看完的视频
				if preLearningNote.Data.StudiedLessonDto.LearnTimeSec != lesson.VideoSec {
					tempMap["watchPoint"] = ""
					tempMap["ev"] = ""
					tempMap["learningTokenId"] = base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(preLearningNote.Data.StudiedLessonDto.ID, 10)))
					tempMap["courseId"] = util.IntToString(videoList.Data.CourseID)
					tempMap["uuid"] = uuid

					// 计算ev参数
					tempMap["recruitId"] = util.IntToString(videoList.Data.RecruitID)
					tempMap["lessonId"] = util.IntToString(lesson.ID)
					tempMap["small"] = util.IntToString(0)
					tempMap["videoId"] = util.IntToString(lesson.VideoID)
					tempMap["chapterId"] = util.IntToString(lesson.ChapterID)
					tempMap["status"] = util.IntToString(0)
					tempMap["playSec"] = util.IntToString(lesson.VideoSec - preLearningNote.Data.StudiedLessonDto.LearnTimeSec) // 用视频的总时长 - 上次播放的可得本次需要播放的总时长
					tempMap["totalSec"] = util.IntToString(preLearningNote.Data.StudiedLessonDto.LearnTimeSec) // 已经缓存（上次播放的总时长）
					tempMap["position"] = util.SecToTim(lesson.VideoSec)

					tempMap["name"] = lesson.Name
					tempMap["chapterName"] = chapter.Name

					if nowTime == 0{
						nowTime = int(time.Now().Unix())
						fmt.Println("nowTime (当前时间) :", nowTime)
					}
					parseInt, _ := strconv.ParseInt(tempMap["playSec"], 10, 64)
					waitTime := math.Ceil(float64(parseInt) / 1.5)
					if isNeed == 1 {
						nowTime += int(waitTime)
					}
					fmt.Println("开始执行", tempMap["name"], "任务")
					fmt.Println("该课程已播放", tempMap["totalSec"], "秒，总时长为：", tempMap["position"], "还需观看：", tempMap["playSec"], "秒")
					if isNeed == 1 {
						formatTimes := time.Unix(int64(nowTime), 0)
						fmt.Println("模拟1.5倍速 需要等候的时间：", waitTime, "秒")
						fmt.Println("预计完成时间为：", formatTimes.String())
						time.Sleep(time.Duration(waitTime) * time.Second)
					}
					saveStudyTime(tempMap)
				}
			}
		}
	}
}
