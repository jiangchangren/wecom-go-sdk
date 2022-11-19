package wework

import (
	"github.com/go-laoji/wecom-go-sdk/internal"
)

type CreateScheduleReq struct {
	Schedule Schedule `json:"schedule"`
	AgentId  int      `json:"agentid"`
}

type Schedule struct {
	Organizer               string              `json:"organizer,omitempty"`
	StartTime               int64               `json:"start_time"`
	EndTime                 int64               `json:"end_time"`
	Attendees               []*ScheduleAttendee `json:"attendees,omitempty"`
	Summary                 string              `json:"summary,omitempty"`
	Description             string              `json:"description,omitempty"`
	Reminders               *ScheduleReminders  `json:"reminders,omitempty"`
	Location                string              `json:"location,omitempty"`
	AllowActiveJoin         int                 `json:"allow_active_join"`
	OnlyOrganizerCreateChat int                 `json:"only_organizer_create_chat"`
	CalId                   string              `json:"cal_id,omitempty"`
	ScheduleId              string              `json:"schedule_id,omitempty"`
}

type ScheduleReminders struct {
	IsRemind              int   `json:"is_remind"`
	RemindBeforeEventSecs int   `json:"remind_before_event_secs"`
	IsRepeat              int   `json:"is_repeat"`
	RepeatType            int   `json:"repeat_type"`
	RepeatUntil           int   `json:"repeat_until"`
	IsCustomRepeat        int   `json:"is_custom_repeat"`
	RepeatInterval        int   `json:"repeat_interval"`
	RepeatDayOfWeek       []int `json:"repeat_day_of_week"`
	RepeatDayOfMonth      []int `json:"repeat_day_of_month"`
	Timezone              int   `json:"timezone"`
}

type CreateScheduleResponse struct {
	internal.BizResponse
	ScheduleId string `json:"schedule_id"`
}

type GetScheduleReq struct {
	ScheduleIdList []string `json:"schedule_id_list"`
}

type GetScheduleReply struct {
	internal.BizResponse
	ScheduleList []*Schedule `json:"schedule_list"`
}

type UpdateScheduleReq struct {
	SkipAttendees int      `json:"skip_attendees"`
	OpMode        int      `json:"op_mode"`
	OpStartTime   int      `json:"op_start_time"`
	Schedule      Schedule `json:"schedule"`
	AgentId       int      `json:"agentid,omitempty"`
}

type DeleteScheduleReq struct {
	ScheduleId  string `json:"schedule_id"`
	OpMode      int    `json:"op_mode"`
	OpStartTime int    `json:"op_start_time"`
}

type ScheduleAttendeesReq struct {
	ScheduleId string              `json:"schedule_id"`
	Attendees  []*ScheduleAttendee `json:"attendees"`
}

type ScheduleAttendee struct {
	Userid string `json:"userid"`
}

// CreateSchedule 创建日程
func (ww weWork) CreateSchedule(corpId uint, req *CreateScheduleReq) (resp CreateScheduleResponse) {
	ww.request(corpId, "/cgi-bin/oa/schedule/add", req, &resp)
	return
}

// GetSchedule 获取日程
func (ww weWork) GetSchedule(corpId uint, req *GetScheduleReq) (resp GetScheduleReply) {
	ww.request(corpId, "/cgi-bin/oa/schedule/get", req, &resp)
	return
}

// UpdateSchedule 更新日程
func (ww weWork) UpdateSchedule(corpId uint, req *UpdateScheduleReq) (resp CreateScheduleResponse) {
	ww.request(corpId, "/cgi-bin/oa/schedule/update", req, &resp)
	return
}

// DeleteSchedule 删除日程
func (ww weWork) DeleteSchedule(corpId uint, req *DeleteScheduleReq) (resp internal.BizResponse) {
	ww.request(corpId, "/cgi-bin/oa/schedule/del", req, &resp)
	return
}

// AddScheduleAttendees 新增日程参与者
func (ww weWork) AddScheduleAttendees(corpId uint, req *ScheduleAttendeesReq) (resp internal.BizResponse) {
	ww.request(corpId, "/cgi-bin/oa/schedule/add_attendees", req, &resp)
	return
}

// DeleteScheduleAttendees 删除日程参与者
func (ww weWork) DeleteScheduleAttendees(corpId uint, req *ScheduleAttendeesReq) (resp internal.BizResponse) {
	ww.request(corpId, "/cgi-bin/oa/schedule/del_attendees", req, &resp)
	return
}
