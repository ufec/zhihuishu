package zhihuishu


type ShareCourseInfoStruct struct {
	Result struct {
		TotalCount     int `json:"totalCount"`
		CourseOpenDtos []struct {
			CourseID        int         `json:"courseId"`
			CourseName      string      `json:"courseName"`
			CourseImg       string      `json:"courseImg"`
			TeacherName     string      `json:"teacherName"`
			LessonName      string      `json:"lessonName"`
			LessonNum       string      `json:"lessonNum"`
			Progress        string      `json:"progress"`
			RecruitID       int         `json:"recruitId"`
			LessonID        interface{} `json:"lessonId"`
			CourseCount     interface{} `json:"courseCount"`
			SchoolName      string      `json:"schoolName"`
			Secret          string      `json:"secret"`
			CourseStartTime int64       `json:"courseStartTime"`
			CourseEndTime   int64       `json:"courseEndTime"`
			Status          int         `json:"status"`
			CourseStatus    int         `json:"courseStatus"`
			CourseType      int         `json:"courseType"`
		} `json:"courseOpenDtos"`
	} `json:"result"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type CourseInfoStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		IsApply          int  `json:"isApply"`
		ClassID          int  `json:"classId"`
		ShowStudyBefore  bool `json:"showStudyBefore"`
		HasExperiment    bool `json:"hasExperiment"`
		SchoolID         int  `json:"schoolId"`
		ScorePublishRule int  `json:"scorePublishRule"`
		RecruitID        int  `json:"recruitId"`
		CourseInfo       struct {
			ID                  interface{} `json:"id"`
			CourseID            int         `json:"courseId"`
			CanJumpChapter      bool        `json:"canJumpChapter"`
			MyuniCourseID       interface{} `json:"myuniCourseId"`
			ExamModelID         interface{} `json:"examModelId"`
			StudyModelID        interface{} `json:"studyModelId"`
			Name                string      `json:"name"`
			EnName              interface{} `json:"enName"`
			OpenTime            interface{} `json:"openTime"`
			Introduction        interface{} `json:"introduction"`
			Hit                 int         `json:"hit"`
			Img                 string      `json:"img"`
			CourseOrg           interface{} `json:"courseOrg"`
			CourseTime          interface{} `json:"courseTime"`
			CommentShow         interface{} `json:"commentShow"`
			State               interface{} `json:"state"`
			Credit              interface{} `json:"credit"`
			CreditPrice         interface{} `json:"creditPrice"`
			IsOuter             interface{} `json:"isOuter"`
			SiteURL             interface{} `json:"siteUrl"`
			MajorSpeaker        interface{} `json:"majorSpeaker"`
			MajorSpeakerName    interface{} `json:"majorSpeakerName"`
			CreditLevel         interface{} `json:"creditLevel"`
			LetterIndex         interface{} `json:"letterIndex"`
			SchoolID            interface{} `json:"schoolId"`
			SchoolName          string      `json:"schoolName"`
			IsDeleted           int         `json:"isDeleted"`
			UpdateTime          interface{} `json:"updateTime"`
			CreateTime          interface{} `json:"createTime"`
			DeletePerson        interface{} `json:"deletePerson"`
			CreatePerson        interface{} `json:"createPerson"`
			UserRange           interface{} `json:"userRange"`
			ExcellentType       interface{} `json:"excellentType"`
			TagString           interface{} `json:"tagString"`
			LabelIds            interface{} `json:"labelIds"`
			TagIds              interface{} `json:"tagIds"`
			TeacherName         interface{} `json:"teacherName"`
			TeacherURL          interface{} `json:"teacherURL"`
			ApprovalSuggestion  interface{} `json:"approvalSuggestion"`
			LabelString         interface{} `json:"labelString"`
			CountSelected       interface{} `json:"countSelected"`
			LabelProxy          interface{} `json:"labelProxy"`
			RecruitStatus       interface{} `json:"recruitStatus"`
			PayStatus           interface{} `json:"payStatus"`
			Paymoney            interface{} `json:"paymoney"`
			RecruitStartTime    interface{} `json:"recruitStartTime"`
			Rnumber             interface{} `json:"rnumber"`
			RecruitID           interface{} `json:"recruitId"`
			RecruitType         interface{} `json:"recruitType"`
			LinkID              interface{} `json:"linkId"`
			MeetCourseState     interface{} `json:"meetCourseState"`
			JobStatus           interface{} `json:"jobStatus"`
			IsShareCourse       int         `json:"isShareCourse"`
			CopyCourseID        interface{} `json:"copyCourseId"`
			CourseMode          int         `json:"courseMode"`
			ApprovalQueryStatus interface{} `json:"approvalQueryStatus"`
			CourseCategory      interface{} `json:"courseCategory"`
			Progress            interface{} `json:"progress"`
			RootCourseID        interface{} `json:"rootCourseId"`
			ObjNum              int         `json:"objNum"`
			SubNum              int         `json:"subNum"`
			DeletelabelIds      interface{} `json:"deletelabelIds"`
			SchoolLogo          interface{} `json:"schoolLogo"`
			CourseOpener        int         `json:"courseOpener"`
			Type                int         `json:"type"`
			IsFree              interface{} `json:"isFree"`
			Price               interface{} `json:"price"`
			MainPeople          interface{} `json:"mainPeople"`
			TurnType            int         `json:"turnType"`
			CourseSpeakers      interface{} `json:"courseSpeakers"`
			PcCourseID          interface{} `json:"pcCourseId"`
			ChapterDisplayMode  int         `json:"chapterDisplayMode"`
			MovieDesc           bool        `json:"movieDesc"`
		} `json:"courseInfo"`
		HasHabbitScore bool   `json:"hasHabbitScore"`
		ClassSchedule  bool   `json:"classSchedule"`
		StudyStatus    string `json:"studyStatus"`
		RunStandard    int    `json:"runStandard"`
	} `json:"data"`
}

type VideoListStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RecruitID        int `json:"recruitId"`
		CourseID         int `json:"courseId"`
		VideoChapterDtos []struct {
			ID                   int    `json:"id"`
			Name                 string `json:"name"`
			OrderNumber          int    `json:"orderNumber"`
			Description          string `json:"description"`
			LearningDay          int    `json:"learningDay"`
			LearningLessonStatus int    `json:"learningLessonStatus"`
			IsPassBadge          int    `json:"isPassBadge"`
			LimitChecked         int    `json:"limitChecked"`
			VideoLessons         []struct {
				ID                   int    `json:"id"`
				Name                 string `json:"name"`
				OrderNumber          int    `json:"orderNumber"`
				Introduction         string `json:"introduction"`
				ChapterID            int    `json:"chapterId"`
				VideoID              int    `json:"videoId"`
				VideoSec             int    `json:"videoSec"`
				IsDeteled            int    `json:"isDeteled"`
				IshaveChildrenLesson int    `json:"ishaveChildrenLesson"`
				IsStudiedLesson      int    `json:"isStudiedLesson"`
			} `json:"videoLessons"`
			IsPass         int `json:"isPass"`
			StudentExamDto struct {
				ID               int         `json:"id"`
				CorrectProgress  interface{} `json:"correctProgress"`
				CreateTime       interface{} `json:"createTime"`
				UpdateTime       interface{} `json:"updateTime"`
				CourseID         interface{} `json:"courseId"`
				ClassID          interface{} `json:"classId"`
				UserID           interface{} `json:"userId"`
				IsDelete         int         `json:"isDelete"`
				OneHundredPont   interface{} `json:"oneHundredPont"`
				GradePont        interface{} `json:"gradePont"`
				LevelPont        interface{} `json:"levelPont"`
				PassPont         interface{} `json:"passPont"`
				IsUpdate         interface{} `json:"isUpdate"`
				RemainingTime    interface{} `json:"remainingTime"`
				State            int         `json:"state"`
				AchieveCount     interface{} `json:"achieveCount"`
				Achieve          interface{} `json:"achieve"`
				CourseSourceType interface{} `json:"courseSourceType"`
				Exam             struct {
					ID        int         `json:"id"`
					IsRecruit int         `json:"isRecruit"`
					LimitTime interface{} `json:"limitTime"`
					ParentID  interface{} `json:"parentId"`
					Type      int         `json:"type"`
					StartDate int64       `json:"startDate"`
					ExamPaper struct {
						Type interface{} `json:"type"`
					} `json:"examPaper"`
					ReviewEndTime interface{} `json:"reviewEndTime"`
					LateSetting   interface{} `json:"lateSetting"`
					LateScore     interface{} `json:"lateScore"`
					TotalScore    interface{} `json:"totalScore"`
				} `json:"exam"`
				Remark        interface{} `json:"remark"`
				PublishDate   interface{} `json:"publishDate"`
				EndDate       interface{} `json:"endDate"`
				Score         interface{} `json:"score"`
				ExamID        interface{} `json:"examId"`
				ExamStartFlag interface{} `json:"examStartFlag"`
				ParentID      interface{} `json:"parentId"`
				ExamURL       string      `json:"examUrl"`
			} `json:"studentExamDto"`
		} `json:"videoChapterDtos"`
	} `json:"data"`
}

type PreLearningNoteStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		StudiedLessonDto struct {
			ID               int64       `json:"id"`
			UserID           interface{} `json:"userId"`
			LastStudyTime    interface{} `json:"lastStudyTime"`
			IsDeteled        interface{} `json:"isDeteled"`
			UpdateTime       interface{} `json:"updateTime"`
			CreateTime       interface{} `json:"createTime"`
			LearnTime        string      `json:"learnTime"`
			LearnTimeSec     int         `json:"learnTimeSec"`
			IsStudiedLesson  int         `json:"isStudiedLesson"`
			StudyTotalTime   int         `json:"studyTotalTime"`
			StudyPercent     interface{} `json:"studyPercent"`
			WatchState       int         `json:"watchState"`
			VideoSize        float64     `json:"videoSize"`
			SourseType       interface{} `json:"sourseType"`
			WatchCount       interface{} `json:"watchCount"`
			PlayTimes        interface{} `json:"playTimes"`
			VideoID          interface{} `json:"videoId"`
			RecruitID        interface{} `json:"recruitId"`
			LessonID         interface{} `json:"lessonId"`
			LessonVideoID    interface{} `json:"lessonVideoId"`
			ChapterID        interface{} `json:"chapterId"`
			PersonalCourseID interface{} `json:"personalCourseId"`
		} `json:"studiedLessonDto"`
	} `json:"data"`
}

type GetEvReqJsonStruct struct {
	RecruitId string `json:"recruitId"`
	LessonId  string `json:"lessonId"`
	Small     string `json:"small"`
	VideoId   string `json:"videoId"`
	ChapterId string `json:"chapterId"`
	Status    string `json:"status"`
	PlaySec   string `json:"playSec"`
	TotalSec  string `json:"totalSec"`
	Position  string `json:"position"`
}

type GetEvResStruct struct {
	EV string `json:"ev"`
}

type SaveStudyTimeStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SubmitSuccess bool `json:"submitSuccess"`
	} `json:"data"`
}